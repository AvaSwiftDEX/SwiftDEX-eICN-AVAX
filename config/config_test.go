package config

import (
	"math/big"
	"os"
	"path/filepath"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	// 创建临时测试文件
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "test_config.yaml")

	// 测试加载不存在的配置文件（应该创建默认配置）
	config, err := LoadConfig(configPath)
	assert.NoError(t, err)
	assert.NotNil(t, config)

	// 验证默认值
	assert.Equal(t, uint16(8080), config.HTTP.Port)
	assert.Equal(t, "127.0.0.1", config.HTTP.Host)
	assert.Equal(t, big.NewInt(1), config.Chain.ID)
	assert.Equal(t, "http://127.0.0.1:8545", config.Chain.HTTPURL)
	assert.Equal(t, "ws://127.0.0.1:8546", config.Chain.WSURL)
	assert.True(t, config.Chain.UseFile)
	assert.Equal(t, common.HexToAddress("0x0000000000000000000000000000000000000000"), config.Chain.Address)

	// 修改配置
	config.HTTP.Port = 9090
	config.Chain.ID = big.NewInt(2)
	config.Chain.Address = common.HexToAddress("0x742d35Cc6634C0532925a3b844Bc454e4438f44e")

	// 保存配置
	err = config.SaveConfig(configPath)
	assert.NoError(t, err)

	// 重新加载配置
	newConfig, err := LoadConfig(configPath)
	assert.NoError(t, err)
	assert.NotNil(t, newConfig)

	// 验证修改后的值
	assert.Equal(t, uint16(9090), newConfig.HTTP.Port)
	assert.Equal(t, big.NewInt(2), newConfig.Chain.ID)
	assert.Equal(t, common.HexToAddress("0x742d35Cc6634C0532925a3b844Bc454e4438f44e"), newConfig.Chain.Address)
}

func TestLoadConfigWithKeyFile(t *testing.T) {
	// 创建临时测试文件夹
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "test_config.yaml")

	// 创建一个测试密钥文件
	keyDir := filepath.Join(tmpDir, "keystore")
	err := os.MkdirAll(keyDir, 0755)
	assert.NoError(t, err)

	// 创建配置
	config := &Config{}
	config.HTTP.Port = 8080
	config.HTTP.Host = "127.0.0.1"
	config.Chain.ID = big.NewInt(1)
	config.Chain.HTTPURL = "http://127.0.0.1:8545"
	config.Chain.WSURL = "ws://127.0.0.1:8546"
	config.Chain.KeyFile = filepath.Join(keyDir, "test_key.json")
	config.Chain.UseFile = false // 设置为 false，因为测试环境中没有实际的密钥文件
	config.Chain.Address = common.HexToAddress("0x0000000000000000000000000000000000000000")

	// 保存配置
	err = config.SaveConfig(configPath)
	assert.NoError(t, err)

	// 加载配置
	loadedConfig, err := LoadConfig(configPath)
	assert.NoError(t, err)
	assert.NotNil(t, loadedConfig)

	// 验证配置
	assert.Equal(t, config.Chain.KeyFile, loadedConfig.Chain.KeyFile)
	assert.Equal(t, config.Chain.UseFile, loadedConfig.Chain.UseFile)
}

func TestInvalidConfigFile(t *testing.T) {
	// 创建临时测试文件
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "invalid_config.yaml")

	// 写入无效的 YAML 内容
	err := os.WriteFile(configPath, []byte("invalid: yaml: content: ], }"), 0644)
	assert.NoError(t, err)

	// 尝试加载无效配置
	_, err = LoadConfig(configPath)
	assert.Error(t, err)
}

func TestConfigSaveError(t *testing.T) {
	// 创建一个只读目录
	tmpDir := t.TempDir()
	readOnlyDir := filepath.Join(tmpDir, "readonly")
	err := os.MkdirAll(readOnlyDir, 0444)
	assert.NoError(t, err)

	// 尝试在只读目录中创建配置文件
	configPath := filepath.Join(readOnlyDir, "config.yaml")
	config := &Config{}

	// 保存配置应该失败
	err = config.SaveConfig(configPath)
	assert.Error(t, err)
}

func TestReadPrivateKey(t *testing.T) {
	// 测试不存在的密钥文件
	_, err := readPrivateKeyFromFile("nonexistent_key.json")
	assert.Error(t, err)

	// 测试无效的密钥文件内容
	tmpDir := t.TempDir()
	invalidKeyPath := filepath.Join(tmpDir, "invalid_key.json")
	err = os.WriteFile(invalidKeyPath, []byte("invalid key content"), 0644)
	assert.NoError(t, err)

	_, err = readPrivateKeyFromFile(invalidKeyPath)
	assert.Error(t, err)
}
