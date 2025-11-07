package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMessageHandler(t *testing.T) {
	// 创建一个模拟的 HTTP 请求
	req, err := http.NewRequest("GET", "/api/message", nil)
	if err != nil {
		t.Fatal(err)
	}

	// 创建一个 ResponseRecorder (一个 http.ResponseWriter 的实现) 来记录响应
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(messageHandler)

	// 调用我们想要测试的处理函数
	handler.ServeHTTP(rr, req)

	// 检查状态码是否是 200 OK
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// 检查 Content-Type 响应头
	expectedContentType := "application/json"
	if contentType := rr.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("handler returned wrong content type: got %v want %v",
			contentType, expectedContentType)
	}

	// 检查响应体内容
	var expectedResponse Message
	expectedResponse.Content = "Hello from Go Backend API!"

	var actualResponse Message
	if err := json.NewDecoder(rr.Body).Decode(&actualResponse); err != nil {
		t.Fatalf("could not decode response body: %v", err)
	}

	if actualResponse.Content != expectedResponse.Content {
		t.Errorf("handler returned unexpected body: got %v want %v",
			actualResponse.Content, expectedResponse.Content)
	}
}
