package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
)

// ITransmitterClient 定义了 Transmitter 客户端的接口
type ITransmitterClient interface {
	CrossReceive(data1, data2 []byte) error
	RegisterEICN(url string, chainID int) error
	SyncHeader(chainID *big.Int, number *big.Int, root common.Hash) error
}

// TransmitterClient 结构体用于与 Transmitter 服务器通信
type TransmitterClient struct {
	baseURL string
}

// 确保 TransmitterClient 实现了 ITransmitterClient 接口
var _ ITransmitterClient = (*TransmitterClient)(nil)

// NewTransmitterClient 创建一个新的 TransmitterClient 实例
func NewTransmitterClient(host string, port string) *TransmitterClient {
	return &TransmitterClient{
		baseURL: fmt.Sprintf("http://%s:%s", host, port),
	}
}

// RequestBody 结构体表示 CrossReceive 请求参数
type RequestBody struct {
	Data1 []byte `json:"data1"`
	Data2 []byte `json:"data2"`
}

// RegisterRequest 结构体表示 registerEICN 请求参数
type RegisterRequest struct {
	URL     string `json:"url"`
	ChainID int    `json:"chainId"`
}

// ResponseBody 结构体表示服务器返回数据
type ResponseBody struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}

// RequestHeader 结构体表示 SyncHeader 请求参数
type RequestHeader struct {
	ChainID *big.Int    `json:"chainId"`
	Number  *big.Int    `json:"number"`
	Root    common.Hash `json:"root"`
}

// CrossReceive 发送跨链数据到服务器
func (c *TransmitterClient) CrossReceive(data1, data2 []byte) error {
	reqBody := RequestBody{
		Data1: data1,
		Data2: data2,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return fmt.Errorf("序列化请求数据失败: %v", err)
	}

	resp, err := http.Post(c.baseURL+"/CrossReceive", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("发送请求失败: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("服务器返回错误状态码 %d: %s", resp.StatusCode, string(body))
	}

	var response ResponseBody
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return fmt.Errorf("解析响应数据失败: %v", err)
	}

	if !response.Success {
		return fmt.Errorf("服务器处理失败: %s", response.Message)
	}

	return nil
}

// RegisterEICN 注册 URL 和 ChainID
func (c *TransmitterClient) RegisterEICN(url string, chainID int) error {
	reqBody := RegisterRequest{
		URL:     url,
		ChainID: chainID,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return fmt.Errorf("序列化请求数据失败: %v", err)
	}

	resp, err := http.Post(c.baseURL+"/registerEICN", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("发送请求失败: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("服务器返回错误状态码 %d: %s", resp.StatusCode, string(body))
	}

	var response ResponseBody
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return fmt.Errorf("解析响应数据失败: %v", err)
	}

	if !response.Success {
		return fmt.Errorf("服务器处理失败: %s", response.Message)
	}

	return nil
}

// SyncHeader 发送区块头同步数据到服务器
func (c *TransmitterClient) SyncHeader(chainID *big.Int, number *big.Int, root common.Hash) error {
	reqBody := RequestHeader{
		ChainID: chainID,
		Number:  number,
		Root:    root,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return fmt.Errorf("序列化请求数据失败: %v", err)
	}

	resp, err := http.Post(c.baseURL+"/SyncHeader", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("发送请求失败: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("服务器返回错误状态码 %d: %s", resp.StatusCode, string(body))
	}

	var response ResponseBody
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return fmt.Errorf("解析响应数据失败: %v", err)
	}

	if !response.Success {
		return fmt.Errorf("服务器处理失败: %s", response.Message)
	}

	return nil
}
