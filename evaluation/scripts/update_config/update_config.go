package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/kimroniny/SuperRunner-eICN-eth2/config"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

type RPCRequest struct {
	JsonRPC string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	ID      int         `json:"id"`
}

type RPCResponse struct {
	Result struct {
		Blockchains []struct {
			Name string `json:"name"`
			ID   string `json:"id"`
		} `json:"blockchains"`
	} `json:"result"`
}

func getCChainURL(pURL string, chainName string) (string, error) {
	rpcReq := RPCRequest{
		JsonRPC: "2.0",
		Method:  "platform.getBlockchains",
		Params:  map[string]interface{}{},
		ID:      1,
	}

	reqBody, err := json.Marshal(rpcReq)
	if err != nil {
		return "", fmt.Errorf("error marshaling request: %v", err)
	}

	resp, err := http.Post(pURL, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return "", fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	var rpcResp RPCResponse
	if err := json.NewDecoder(resp.Body).Decode(&rpcResp); err != nil {
		return "", fmt.Errorf("error decoding response: %v", err)
	}

	for _, chain := range rpcResp.Result.Blockchains {
		if chain.Name == chainName {
			return chain.ID, nil
		}
	}

	return "", fmt.Errorf("chain %s not found", chainName)
}

func updateChainURLs(pHost string, pPort int, cHost string, cPort int, cName string, configPath string) error {

	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		return fmt.Errorf("error loading config: %v", err)
	}

	pURL := fmt.Sprintf("http://%s:%d/ext/bc/P", pHost, pPort)

	// Get chain ID and update URLs
	chainID, err := getCChainURL(pURL, cName)
	if err != nil {
		return fmt.Errorf("error getting chain URL: %v", err)
	}

	cfg.Chain.HTTPURL = fmt.Sprintf("http://%s:%d/ext/bc/%s/rpc", cHost, cPort, chainID)
	cfg.Chain.WSURL = fmt.Sprintf("ws://%s:%d/ext/bc/%s/ws", cHost, cPort, chainID)

	// Write back to YAML file
	newData, err := yaml.Marshal(&cfg)
	if err != nil {
		return fmt.Errorf("error marshaling config: %v", err)
	}

	if err := os.WriteFile(configPath, newData, 0644); err != nil {
		return fmt.Errorf("error writing config file: %v", err)
	}

	return nil
}

func main() {
	var pHost, cHost, cName, configPath string
	var pPort, cPort int

	var rootCmd = &cobra.Command{
		Use:   "update_config",
		Short: "Update chain configuration URLs",
		Long:  "A CLI tool to update chain configuration URLs by fetching chain ID from platform API",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cPort == 0 || cName == "" || configPath == "" {
				return fmt.Errorf("required flags: c_port, c_name, and config_path must be provided")
			}

			absConfigPath, err := filepath.Abs(configPath)
			if err != nil {
				return fmt.Errorf("error getting absolute path: %v", err)
			}

			return updateChainURLs(pHost, pPort, cHost, cPort, cName, absConfigPath)
		},
	}

	rootCmd.Flags().StringVar(&pHost, "p_host", "127.0.0.1", "Platform host")
	rootCmd.Flags().IntVar(&pPort, "p_port", 60010, "Platform port")
	rootCmd.Flags().StringVar(&cHost, "c_host", "127.0.0.1", "Chain host")
	rootCmd.Flags().IntVar(&cPort, "c_port", 0, "Chain port")
	rootCmd.Flags().StringVar(&cName, "c_name", "", "Chain name")
	rootCmd.Flags().StringVar(&configPath, "config_path", "", "Config file path")

	// Mark required flags
	rootCmd.MarkFlagRequired("c_port")
	rootCmd.MarkFlagRequired("c_name")
	rootCmd.MarkFlagRequired("config_path")

	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
