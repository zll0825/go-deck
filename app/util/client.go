package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

var transport *http.Transport

type Service struct {
	Url         string
	Host        string
	Method      string
	ContentType string
	Client      *http.Client
	FormatFunc  string
	LoginIgnore bool
	ServiceName string
}

func createTransport() *http.Transport {
	if transport != nil {
		return transport
	}

	transport := &http.Transport{
		DisableKeepAlives:   false,
		MaxIdleConnsPerHost: 10,
	}
	return transport
}

// 新建http client
func New(timeout int) *http.Client {
	client := &http.Client{
		Transport: createTransport(),
		Timeout:   time.Duration(timeout) * time.Millisecond,
	}
	return client
}

/**
 * 简单的post方法
 * params为post的参数列表
 * f 为约定的返回值类型，结果可以直接读取
 * 最终会返回一个string类型的返回值供调用放自己使用
 */
func DoPost(ctx *gin.Context, service *Service, params map[string]interface{}, f interface{}) (string, error) {
	// 参数注入
	pByte, _ := json.Marshal(params)

	request, _ := http.NewRequest(service.Method, service.Url, bytes.NewBuffer(pByte))
	request.Header.Set("Content-Type", "application/json")
	// do call
	retBody, err := DoCall(ctx, service, request)

	if err != nil {
		return "", err
	}

	// json 处理
	err = json.Unmarshal(retBody, &f)
	if err != nil {
		// todo log
		return "", err
	}

	return string(retBody), err
}

/**
 * 简单的get方法
 * params为post的参数列表
 * f 为约定的返回值类型，结果可以直接读取
 * 最终会返回一个string类型的返回值供调用放自己使用
 */
func DoGet(ctx *gin.Context, service *Service, params map[string]interface{}, f interface{}) (string, error) {
	// 参数注入
	url, err := GenerateUri(service.Url, params)
	if err != nil {
		// todo log
		return "", err
	}
	var br io.Reader
	request, _ := http.NewRequest(service.Method, url, br)
	request.Header.Set("Content-Type", "application/json")
	// do call
	retBody, err := DoCall(ctx, service, request)

	if err != nil {
		return "", err
	}

	// json 处理
	err = json.Unmarshal(retBody, &f)
	if err != nil {
		// todo log
		return "", err
	}

	return string(retBody), err
}

/**
 * do http request
 */
func DoCall(ctx *gin.Context, s *Service, req *http.Request) ([]byte, error) {
	_ = ctx.Request.ParseMultipartForm(32 << 20)

	resp, err := s.Client.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	} else if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Response status: %s ", resp.Status)
	}

	return respBody, nil
}
