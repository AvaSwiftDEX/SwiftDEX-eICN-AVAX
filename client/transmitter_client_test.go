package client

import (
	"encoding/json"
	"math/big"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ethereum/go-ethereum/common"
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

	// 创建存储和客户端
	storage := make(map[string]string)
	chainID := big.NewInt(1)
	storage[chainID.String()] = server.URL
	client := NewTransmitterClient(storage)

	// 测试数据
	data1 := []byte("test data 1")
	data2 := []byte("test data 2")

	// 执行测试
	err := client.CrossReceive(chainID, data1, data2)
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

	// 创建存储和客户端
	storage := make(map[string]string)
	client := NewTransmitterClient(storage)

	// 测试数据
	url := "http://example.com"
	chainID := big.NewInt(1)

	// 执行测试
	err := client.RegisterEICN(url, chainID, server.URL)
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

	// 创建存储和客户端
	storage := make(map[string]string)
	chainID := big.NewInt(1)
	storage[chainID.String()] = server.URL
	client := NewTransmitterClient(storage)

	// 测试数据
	number := big.NewInt(100)
	root := common.HexToHash("0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef")

	// 执行测试
	err := client.SyncHeader(chainID, number, root)
	if err != nil {
		t.Errorf("SyncHeader 失败: %v", err)
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

	// 创建存储和客户端
	storage := make(map[string]string)
	chainID := big.NewInt(1)
	storage[chainID.String()] = server.URL
	client := NewTransmitterClient(storage)

	// 测试 CrossReceive 错误处理
	err := client.CrossReceive(chainID, []byte("test"), []byte("test"))
	if err == nil {
		t.Error("期望 CrossReceive 返回错误，但得到 nil")
	}

	// 测试 RegisterEICN 错误处理
	err = client.RegisterEICN("http://example.com", chainID, server.URL)
	if err == nil {
		t.Error("期望 RegisterEICN 返回错误，但得到 nil")
	}

	// 测试 SyncHeader 错误处理
	root := common.HexToHash("0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef")
	err = client.SyncHeader(chainID, big.NewInt(100), root)
	if err == nil {
		t.Error("期望 SyncHeader 返回错误，但得到 nil")
	}
}
