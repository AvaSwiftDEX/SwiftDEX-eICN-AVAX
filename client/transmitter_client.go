package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/kimroniny/SuperRunner-eICN-eth2/logger"
	"github.com/sirupsen/logrus"
)

// ITransmitterClient 定义了 Transmitter 客户端的接口
type ITransmitterClient interface {
	CrossReceive(chainId *big.Int, data1, data2 []byte) error
	RegisterEICN(url string, chainID *big.Int, targetServerURL string) error
	SyncHeader(chainID *big.Int, number *big.Int, root common.Hash) error
}

// TransmitterClient 结构体用于与 Transmitter 服务器通信
type TransmitterClient struct {
	storage map[string]string
	log     *logrus.Entry
}

// 确保 TransmitterClient 实现了 ITransmitterClient 接口
var _ ITransmitterClient = (*TransmitterClient)(nil)

// NewTransmitterClient 创建一个新的 TransmitterClient 实例
func NewTransmitterClient(storage map[string]string) *TransmitterClient {
	if logger.GetLogger() == nil {
		logger.InitLogger("")
	}
	return &TransmitterClient{
		storage: storage,
		log:     logger.NewComponent("TransmitterClient"),
	}
}

// RequestBody 结构体表示 CrossReceive 请求参数
type RequestBody struct {
	Data1 []byte `json:"data1"`
	Data2 []byte `json:"data2"`
}

// RegisterRequest 结构体表示 registerEICN 请求参数
type RegisterRequest struct {
	URL     string   `json:"url"`
	ChainID *big.Int `json:"chainId"`
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
func (c *TransmitterClient) CrossReceive(chainId *big.Int, data1, data2 []byte) error {
	var targetServerURL string
	if _url, ok := c.storage[chainId.String()]; !ok {
		c.log.Info(fmt.Sprintf("未找到链(#%d)的 URL", chainId))
		return nil
	} else {
		targetServerURL = _url
	}

	reqBody := RequestBody{
		Data1: data1,
		Data2: data2,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return fmt.Errorf("序列化请求数据失败: %v", err)
	}

	c.log.WithFields(logrus.Fields{
		"chainID": chainId,
		"url":     targetServerURL,
	}).Debug("call the counterparty eICN server's CrossReceive")

	resp, err := http.Post(targetServerURL+"/CrossReceive", "application/json", bytes.NewBuffer(jsonData))
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
func (c *TransmitterClient) RegisterEICN(url string, chainID *big.Int, targetServerURL string) error {
	c.log.WithFields(logrus.Fields{
		"url":     url,
		"chainID": chainID,
	}).Info("regist EICN")

	reqBody := RegisterRequest{
		URL:     url,
		ChainID: chainID,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return fmt.Errorf("序列化请求数据失败: %v", err)
	}

	resp, err := http.Post(targetServerURL+"/registerEICN", "application/json", bytes.NewBuffer(jsonData))
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

	wg := sync.WaitGroup{}
	for targetChainID, url := range c.storage {
		wg.Add(1)
		go func(targetChainID string, url string) {
			defer wg.Done()
			err := c.syncHeaderToSingleChain(url, jsonData)
			if err != nil {
				c.log.Warn(fmt.Sprintf("transmit header to chain(#%s)'s eICN server, while error: %v", targetChainID, err))
				return
			}
			c.log.WithFields(logrus.Fields{
				"method":     "SyncHeader",
				"chanID":     chainID,
				"blockNum":   number,
				"headerRoot": root.Hex(),
			}).Debug(fmt.Sprintf("transmit header to chain(#%s)'s eICN server success", targetChainID))
		}(targetChainID, url)
	}
	wg.Wait()

	return nil
}

func (c *TransmitterClient) syncHeaderToSingleChain(targetServerURL string, jsonData []byte) error {
	resp, err := http.Post(targetServerURL+"/SyncHeader", "application/json", bytes.NewBuffer(jsonData))
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
