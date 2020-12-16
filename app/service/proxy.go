package service

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-deck/app/global"
	"go-deck/app/response"
	"go-deck/app/util"
	"go.uber.org/zap"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
)

type Proxy struct {
	c    *gin.Context
	up   *util.Service
	req  *UpstreamRequest
	resp *UpstreamResponse
}

type UpstreamResponse struct {
	// handler func(*gin.Context, *http.Response) error
	resp *http.Response
}

type UpstreamRequest struct {
	query        url.Values
	body         string
	req          *http.Request
	fileWriter   *multipart.Writer
	queryHandler func(*gin.Context, map[string]interface{}, url.Values) error
	reqHandler   func(*gin.Context, *http.Request) error
	bodyHandler  func(*gin.Context, map[string]interface{}, *UpstreamRequest) error
}

func NewProxy(c *gin.Context, s *util.Service) *Proxy {
	return &Proxy{
		c:    c,
		up:   s,
		req:  &UpstreamRequest{},
		resp: &UpstreamResponse{},
	}
}

func (p *Proxy) Run() {
	if !p.parseRequest() {
		return
	}

	if !p.doCall() {
		return
	}

	if !p.buildResponse() {
		return
	}
}

// 格式化request
func (p *Proxy) parseRequest() bool {
	ctx := p.c
	service := p.up
	req := p.req

	var ret bool

	req.query = ctx.Request.URL.Query()
	requestUrl := service.Url

	switch ctx.Request.Method {
	case "POST":
		body, _ := ioutil.ReadAll(ctx.Request.Body)
		ctx.Set("form_body", string(body))
		// json body
		jBody := make(map[string]interface{})
		// 构造转发的body
		retBody := make(map[string]interface{})
		if len(body) > 0 {
			err := json.Unmarshal(body, &jBody)
			if err != nil {
				response.FailWithMessage(ctx, "form body illegal.")
				ctx.Abort()
				return false
			}
		}

		for k, v := range jBody {
			retBody[k] = v
		}

		buf := bytes.NewBuffer(util.GeneratePostBody(retBody, service.ContentType))
		req.body = buf.String()

		req.req, _ = http.NewRequest(service.Method, requestUrl, buf)
		req.req.Header.Set("Content-Type", "application/json; charset=utf-8")
	case "GET":
		var br io.Reader
		req.req, _ = http.NewRequest(ctx.Request.Method, requestUrl, br)
	default:
		response.FailWithMessage(ctx, "build request error")
		ctx.Abort()
		return false
	}
	ret = p.requestHeader()
	global.Logger.Info("proxy request info",
		zap.String("HOST", req.req.URL.Host),
		zap.String("PATH", req.req.URL.Path),
		zap.String("QUERY", req.req.URL.RawQuery))

	return ret
}

// 执行请求
func (p *Proxy) doCall() bool {
	ctx := p.c
	service := p.up
	ureq := p.req
	uresp := p.resp
	var err error

	// call
	uresp.resp, err = service.Client.Do(ureq.req)

	if err != nil {
		// 错误信息返回
		global.Logger.Error("Upstream request error", zap.Any("err", err.Error()))
		response.FailWithMessage(ctx, "Upstream request error.")

		if ureq.fileWriter != nil {
			_ = ureq.fileWriter.Close()
		}

		p.c.Abort()
		return false
	}

	return true
}

// 格式化返回值
func (p *Proxy) buildResponse() bool {
	ctx := p.c
	ureq := p.req
	uresp := p.resp

	if ureq.fileWriter != nil {
		_ = ureq.fileWriter.Close()
	}

	if uresp.resp != nil {
		defer uresp.resp.Body.Close()
	}

	// Check that the server actually sent compressed data
	var reader io.ReadCloser
	switch uresp.resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, _ = gzip.NewReader(uresp.resp.Body)
		defer reader.Close()
	default:
		reader = uresp.resp.Body
	}
	respBody, err := ioutil.ReadAll(reader)

	global.Logger.Info("proxy response",
		zap.String("status", uresp.resp.Status),
		zap.String("res", string(respBody)))

	ctx.Set("upstream_http_code", uresp.resp.StatusCode)

	if err != nil {
		global.Logger.Error("build Response error",
			zap.String("respBody", string(respBody)),
			zap.Any("err", err.Error()))
		response.FailWithMessage(ctx, fmt.Sprintf("Upstream response error[%s]", string(respBody)))
		return false
	}

	var jsonRet map[string]interface{}
	err = json.Unmarshal(respBody, &jsonRet)
	if err != nil {
		response.FailWithMessage(ctx, fmt.Sprintf("Upstream response error[%s]", string(respBody)))
		return false
	}

	response.ProxyResponse(ctx, jsonRet)

	return true
}

func (p *Proxy) requestHeader() bool {
	request := p.req.req
	service := p.up
	if service.LoginIgnore {
		return true
	}

	// copy header
	for k, v := range p.c.Request.Header {
		request.Header.Set(k, v[0])
	}

	// add user info
	//userInfo, exist := p.c.Get(config.CtxUserInfoKey)
	//if val, ok := userInfo.(string); exist && ok {
	//	request.Header.Set("X-UserInfo", val)
	//}
	return true
}
