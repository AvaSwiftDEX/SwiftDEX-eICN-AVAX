package config

import (
	"crypto/ecdsa"
	"encoding/hex"
	"math/big"
	"os"
	"path/filepath"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"gopkg.in/yaml.v3"
)

// Config 定义配置结构
type Config struct {
	HTTP struct {
		Port    uint16 `yaml:"port"`
		Host    string `yaml:"host"`
		BaseURL string `yaml:"base_url"`
	} `yaml:"http"`

	Chain struct {
		ID      *big.Int       `yaml:"id"`
		HTTPURL string         `yaml:"http_url"`
		WSURL   string         `yaml:"ws_url"`
		KeyFile string         `yaml:"key_file"`
		KeyHex  string         `yaml:"key_hex"`
		UseFile bool           `yaml:"use_file"`
		Address common.Address `yaml:"address,omitempty"`
	} `yaml:"chain"`
}

func readPrivateKey(keyFile string) (*ecdsa.PrivateKey, error) {
	key, err := os.ReadFile(keyFile)
	if err != nil {
		return nil, err
	}
	account, err := keystore.DecryptKey(key, "")
	if err != nil {
		return nil, err
	}
	return account.PrivateKey, nil
}

// LoadConfig 从文件加载配置
func LoadConfig(filename string) (*Config, error) {
	// 确保配置文件目录存在
	if err := os.MkdirAll(filepath.Dir(filename), 0755); err != nil {
		return nil, err
	}

	// 读取配置文件
	data, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			// 如果文件不存在，创建默认配置
			return createDefaultConfig(filename)
		}
		return nil, err
	}

	config := &Config{}
	if err := yaml.Unmarshal(data, config); err != nil {
		return nil, err
	}

	// update privateKeyHex if useFile is true
	if config.Chain.UseFile {
		privateKey, err := readPrivateKey(config.Chain.KeyFile)
		if err != nil {
			return nil, err
		}
		config.Chain.KeyHex = hex.EncodeToString(crypto.FromECDSA(privateKey))[2:]
	}
	return config, nil
}

// SaveConfig 保存配置到文件
func (c *Config) SaveConfig(filename string) error {
	data, err := yaml.Marshal(c)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}

// createDefaultConfig 创建默认配置
func createDefaultConfig(filename string) (*Config, error) {
	config := &Config{}

	// 设置默认值
	config.HTTP.Port = 8080
	config.HTTP.Host = "127.0.0.1"
	config.HTTP.BaseURL = "http://127.0.0.1:8080"

	config.Chain.ID = big.NewInt(1)
	config.Chain.HTTPURL = "http://127.0.0.1:8545"
	config.Chain.WSURL = "ws://127.0.0.1:8546"
	config.Chain.KeyFile = "node/keystore/UTC--2025-02-08T06-12-23.376721660Z--a8a410a56f93e14fb5a71f5968958851915b6909"
	config.Chain.KeyHex = "c45ba5d6de0e502aefd23c98b40a2c9018e2e0286dde4fdb542ded619cefc8bd"
	config.Chain.UseFile = true
	config.Chain.Address = common.HexToAddress("0x0000000000000000000000000000000000000000")

	// 保存默认配置到文件
	if err := config.SaveConfig(filename); err != nil {
		return nil, err
	}

	return config, nil
}
