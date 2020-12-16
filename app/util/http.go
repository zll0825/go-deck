package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
)

func GenerateUri(uri string, data map[string]interface{}) (string, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return "", err
	}
	query := u.Query()
	for k, v := range data {
		query.Set(k, fmt.Sprint(v))
	}
	u.RawQuery = query.Encode()
	return u.String(), nil
}

// 构造&分隔的param字符串
func GenerateParam(data map[string]interface{}) (string, error) {
	u, err := url.Parse("")
	if err != nil {
		return "", err
	}
	query := u.Query()
	for k, v := range data {
		query.Set(k, fmt.Sprint(v))
	}
	return query.Encode(), nil
}

// 根据content_type构造post body
func GeneratePostBody(body map[string]interface{}, contentType string) []byte {
	bodyStr := ""
	switch contentType {
	case "application/x-www-form-urlencoded":
		bodyStr, _ = GenerateParam(body)
	case "application/json":
		fallthrough
	default:
		bf := bytes.NewBuffer([]byte{})
		jsonEncoder := json.NewEncoder(bf)
		jsonEncoder.SetEscapeHTML(false)
		_ = jsonEncoder.Encode(body)
		bodyStr = bf.String()
	}

	return []byte(bodyStr)
}
