package sangfor_ad

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// 创建一个跳过 HTTPS 校验的 HTTP 客户端
var insecureHttpClient = &http.Client{
	Timeout: 10 * time.Second,
	Transport: &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	},
}

// 公共请求方法
func requestGet[T any](reqURL string, headers map[string]string) (*T, error) {
	// 构造请求
	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return nil, fmt.Errorf("构建Sangfor AD GET 请求失败: %w", err)
	}

	// 设置额外 headers（如 Authorization 等）
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// 发送请求
	resp, err := insecureHttpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Sangfor AD GET 请求失败: %w", err)
	}
	defer resp.Body.Close()

	// 读取响应体
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取Sangfor AD GET 请求失败: %w", err)
	}

	if resp.StatusCode == 404 {
		return nil, fmt.Errorf("Sangfor AD GET 请求成功,但返回资源为: Not Found")
	}

	// 检查 HTTP 状态码
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("Sangfor AD GET 请求失败, 状态吗为: %d, 响应数据为: %s", resp.StatusCode, string(respBody))
	}

	// 解析 JSON 到泛型 T
	var result T
	if err = json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("Sangfor AD GET 接口JSON 解析失败: %w", err)
	}

	return &result, nil
}

func requestPost[T any](reqURL, method string, params any, headers map[string]string) (*T, error) {
	jsonData, _ := json.Marshal(params)
	// 构造请求
	body := bytes.NewBuffer(jsonData)
	req, err := http.NewRequest(method, reqURL, body)
	if err != nil {
		return nil, fmt.Errorf("构建Sangfor AD %s 请求失败: %w", method, err)
	}

	// 设置 Content-Type
	req.Header.Set("Content-Type", "application/json")

	// 设置额外 headers（如 Authorization 等）
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// 发送请求
	resp, err := insecureHttpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Sangfor AD %s 请求失败: %w", method, err)
	}
	defer resp.Body.Close()

	// 读取响应体
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取Sangfor AD %s 请求响应: %w", method, err)
	}

	// 检查 HTTP 状态码
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("Sangfor AD %s 请求失败, 状态吗为: %d, 响应数据为: %s", method, resp.StatusCode, string(respBody))
	}

	var result T
	if err = json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("Sangfor AD %s 接口JSON 解析失败: %w", method, err)
	}

	return &result, nil
}
