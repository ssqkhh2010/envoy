package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"time"
)

func attemptRequestWithFallback() (response string, success bool) {
	// 定义一个公共变量来存储URL
	targetURL := "http://scserver-service:5000/hello"
	// targetURL = "http://192.168.159.130:5000/hello"

	// 自定义Transport关闭HTTP/2并设置超时
	transport := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   5 * time.Second, // 连接超时
			KeepAlive: 30 * time.Second,
		}).DialContext,
		ForceAttemptHTTP2:     false, // 关闭HTTP/2
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	// 设置带有超时的客户端
	client := &http.Client{
		Timeout:   10 * time.Second, // 整体请求超时
		Transport: transport,
	}

	// 尝试三次普通请求
	for i := 0; i < 3; i++ {
		fmt.Printf("========i:%d\n", i)
		req, err := http.NewRequest("GET", targetURL, nil)
		if err != nil {
			log.Printf("Failed to create request: %v", err)
			continue // 如果创建请求失败，则尝试下一次
		}

		resp, err := client.Do(req)
		if err != nil {
			log.Printf("Request failed: %v", err)
			continue // 如果请求失败，则尝试下一次
		}

		defer resp.Body.Close()

		// 如果请求成功，则返回响应内容
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Failed to read response body: %v", err)
			continue // 如果读取响应体失败，则尝试下一次
		}
		// 检查HTTP状态码，如果请求失败则重试
		if resp.StatusCode >= 300 {
			log.Printf("Request failed with status code: %d, retrying...", resp.StatusCode)
			continue
		}

		return string(body), true // 成功获取响应
	}

	// 所有尝试都失败，执行带有特殊请求头的请求
	req, err := http.NewRequest("GET", targetURL, nil)
	if err != nil {
		log.Printf("Failed to create request with fallback header: %v", err)
		return "Internal server error", false
	}
	fmt.Println("========i:4")
	// 添加特定的请求头
	req.Header.Add("dc", "f")

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Fallback request failed: %v", err)
		return "Internal server error", false
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Failed to read fallback response body: %v", err)
		return "Internal server error", false
	}
	// 检查HTTP状态码，如果请求失败则重试
	if resp.StatusCode >= 300 {
		log.Printf("Request failed with status code: %d, retrying...", resp.StatusCode)
	}
	return string(body), true // 成功获取带有特殊请求头的响应
}

func main() {
	http.HandleFunc("/send", func(w http.ResponseWriter, r *http.Request) {
		response, success := attemptRequestWithFallback()
		if !success {
			http.Error(w, response, http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "Response: %s", response)
	})
	log.Println("Starting server on :8184")
	if err := http.ListenAndServe(":8184", nil); err != nil {
		log.Fatal(err)
	}
}
