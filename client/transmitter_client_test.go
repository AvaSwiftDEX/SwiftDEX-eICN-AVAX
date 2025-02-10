package client

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTransmitterClient_CrossReceive(t *testing.T) {
	// 创建测试服务器
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 验证请求方法和路径
		if r.Method != http.MethodPost {
			t.Errorf("期望 POST 请求，得到 %s", r.Method)
		}
		if r.URL.Path != "/CrossReceive" {
			t.Errorf("期望路径 /CrossReceive，得到 %s", r.URL.Path)
		}

		// 返回成功响应
		response := ResponseBody{Success: true}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	// 创建客户端
	client := &TransmitterClient{baseURL: server.URL}

	// 测试数据
	data1 := []byte("test data 1")
	data2 := []byte("test data 2")

	// 执行测试
	err := client.CrossReceive(data1, data2)
	if err != nil {
		t.Errorf("CrossReceive 失败: %v", err)
	}
}

func TestTransmitterClient_RegisterEICN(t *testing.T) {
	// 创建测试服务器
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 验证请求方法和路径
		if r.Method != http.MethodPost {
			t.Errorf("期望 POST 请求，得到 %s", r.Method)
		}
		if r.URL.Path != "/registerEICN" {
			t.Errorf("期望路径 /registerEICN，得到 %s", r.URL.Path)
		}

		// 返回成功响应
		response := ResponseBody{Success: true}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	// 创建客户端
	client := &TransmitterClient{baseURL: server.URL}

	// 测试数据
	url := "http://example.com"
	chainID := 1

	// 执行测试
	err := client.RegisterEICN(url, chainID)
	if err != nil {
		t.Errorf("RegisterEICN 失败: %v", err)
	}
}

func TestTransmitterClient_SyncHeader(t *testing.T) {
	// 创建测试服务器
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 验证请求方法和路径
		if r.Method != http.MethodPost {
			t.Errorf("期望 POST 请求，得到 %s", r.Method)
		}
		if r.URL.Path != "/SyncHeader" {
			t.Errorf("期望路径 /SyncHeader，得到 %s", r.URL.Path)
		}

		// 返回成功响应
		response := ResponseBody{Success: true}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	// 创建客户端
	client := &TransmitterClient{baseURL: server.URL}

	// 测试数据
	chainID := 1
	number := uint64(100)
	var root [32]byte
	copy(root[:], []byte("test root data"))

	// 执行测试
	err := client.SyncHeader(chainID, number, root)
	if err != nil {
		t.Errorf("SyncHeader 失败: %v", err)
	}
}

func TestNewTransmitterClient(t *testing.T) {
	host := "localhost"
	port := "8080"
	client := NewTransmitterClient(host, port)

	expectedURL := "http://localhost:8080"
	if client.baseURL != expectedURL {
		t.Errorf("期望 baseURL 为 %s，得到 %s", expectedURL, client.baseURL)
	}
}

func TestTransmitterClient_ErrorHandling(t *testing.T) {
	// 创建返回错误的测试服务器
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		response := ResponseBody{
			Success: false,
			Message: "内部服务器错误",
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	// 创建客户端
	client := &TransmitterClient{baseURL: server.URL}

	// 测试 CrossReceive 错误处理
	err := client.CrossReceive([]byte("test"), []byte("test"))
	if err == nil {
		t.Error("期望 CrossReceive 返回错误，但得到 nil")
	}

	// 测试 RegisterEICN 错误处理
	err = client.RegisterEICN("http://example.com", 1)
	if err == nil {
		t.Error("期望 RegisterEICN 返回错误，但得到 nil")
	}

	// 测试 SyncHeader 错误处理
	var root [32]byte
	err = client.SyncHeader(1, 100, root)
	if err == nil {
		t.Error("期望 SyncHeader 返回错误，但得到 nil")
	}
}
