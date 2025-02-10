package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/kimroniny/SuperRunner-eICN-eth2/sdk"
)

// Transmitter 结构体定义 HTTP 服务器
type Transmitter struct {
	port        string
	storage     map[string]int
	mutex       sync.Mutex
	wg          *sync.WaitGroup
	contractSDK *sdk.ContractSDK
}

// NewTransmitter 创建一个新的 Transmitter 实例
func NewTransmitter(port string, wg *sync.WaitGroup, contractSDK *sdk.ContractSDK) *Transmitter {
	return &Transmitter{
		port:        port,
		storage:     make(map[string]int),
		wg:          wg,
		contractSDK: contractSDK,
	}
}

type RequestHeader struct {
	ChainID int    `json:"chainId"`
	Number  uint64 `json:"number"`
	Root    [32]byte `json:"root"`
}

// RequestBody 结构体表示请求参数
type RequestBody struct {
	Data1 []byte `json:"data1"`
	Data2 []byte `json:"data2"`
}

// RegisterRequest 结构体表示 registerEICN 接口的请求参数
type RegisterRequest struct {
	URL     string `json:"url"`
	ChainID int    `json:"chainId"`
}

// ResponseBody 结构体表示返回数据
type ResponseBody struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}

func (t *Transmitter) SyncHeader(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// 读取请求体
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// 解析 JSON
	var req RequestHeader
	err = json.Unmarshal(body, &req)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// 将数据发送到 ContractSDK
	t.contractSDK.TransmitterSyncHeader(req.ChainID, req.Number, req.Root)

	// 返回成功响应
	response := ResponseBody{Success: true}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

// CrossReceive 处理 HTTP 请求，并将数据传输到 ContractSDK
func (t *Transmitter) CrossReceive(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// 读取请求体
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// 解析 JSON
	var req RequestBody
	err = json.Unmarshal(body, &req)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// 将数据发送到 ContractSDK
	t.contractSDK.TransmitterCrossReceive(req.Data1, req.Data2)

	// 返回成功响应
	response := ResponseBody{Success: true}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

// RegisterEICN 处理注册 URL 和 ChainID
func (t *Transmitter) RegisterEICN(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// 读取请求体
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// 解析 JSON
	var req RegisterRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// 存储数据
	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.storage[req.URL] = req.ChainID

	// 返回成功响应
	response := ResponseBody{Success: true}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

// StartServer 启动 HTTP 服务器（协程方式）
func (t *Transmitter) StartServer() {
	defer t.wg.Done()
	http.HandleFunc("/CrossReceive", t.CrossReceive)
	http.HandleFunc("/registerEICN", t.RegisterEICN)
	fmt.Printf("Server is running on port %s...\n", t.port)
	if err := http.ListenAndServe(":"+t.port, nil); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}
