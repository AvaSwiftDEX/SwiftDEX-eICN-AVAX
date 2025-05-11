package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"

	"github.com/kimroniny/SuperRunner-eICN-eth2/config"
	"github.com/kimroniny/SuperRunner-eICN-eth2/scripts/scripts"
	"github.com/spf13/cobra"
)

var (
	// 全局标志
	cfgFile string

	// cross-send 命令的标志
	chainIDsStr string
	value       string

	// regist-eICN 命令的标志
	targetConfigFile string
)

// parseChainIDs 解析链ID参数字符串，格式如 "1,2,3"
func parseChainIDs(chainIDsStr string) ([]*big.Int, error) {
	if chainIDsStr == "" {
		return nil, nil
	}

	idStrs := strings.Split(chainIDsStr, ",")
	chainIDs := make([]*big.Int, len(idStrs))

	for i, idStr := range idStrs {
		id := new(big.Int)
		_, success := id.SetString(idStr, 10)
		if !success {
			return nil, fmt.Errorf("invalid chain ID: %s", idStr)
		}
		chainIDs[i] = id
	}

	return chainIDs, nil
}

func main() {
	rootCmd := &cobra.Command{
		Use:   "sr2pc",
		Short: "SR2PC is a tool for cross-chain operations",
		Long:  `SR2PC provides functionality for deploying contracts and sending cross-chain transactions.`,
	}

	// 全局标志
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "config.yaml", "config file path")

	// Deploy 命令
	deployCmd := &cobra.Command{
		Use:   "deploy",
		Short: "Deploy the SR2PC contract",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := config.LoadConfig(cfgFile)
			if err != nil {
				return fmt.Errorf("failed to load config: %v", err)
			}

			ctx := context.Background()
			address, err := scripts.Deploy(ctx, cfg)
			if err != nil {
				return fmt.Errorf("deploy failed: %v", err)
			}
			cfg.Chain.Address = address
			cfg.SaveConfig(cfgFile)

			fmt.Printf("Successfully deployed contract at: %s\n", address.Hex())
			return nil
		},
	}

	// DeployLibraries 命令
	deployLibrariesCmd := &cobra.Command{
		Use:   "deploy-lib-Filter",
		Short: "Deploy the library Filter",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := config.LoadConfig(cfgFile)
			if err != nil {
				return fmt.Errorf("failed to load config: %v", err)
			}
			ctx := context.Background()
			address, err := scripts.DeployLibraries(ctx, cfg)
			if err != nil {
				return fmt.Errorf("deploy libraries failed: %v", err)
			}

			_ = os.RemoveAll("SR2PC/lib")

			filePath := "SR2PC/lib"
			content := []byte(fmt.Sprintf("../SuperRunner-Contracts/contracts/2pc-master/lib/Filter.sol:Filter=%s", address.Hex()))
			fmt.Println("content:", string(content))
			if err := os.WriteFile(filePath, content, 0644); err != nil {
				return fmt.Errorf("failed to write to file: %v", err)
			}

			fmt.Printf("Successfully deployed library Filter at: %s\n", address.Hex())
			return nil
		},
	}

	// Cross-send 命令
	crossSendCmd := &cobra.Command{
		Use:   "cross-send",
		Short: "Execute a cross-chain send operation",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := config.LoadConfig(cfgFile)
			if err != nil {
				return fmt.Errorf("failed to load config: %v", err)
			}

			// 解析 chain IDs
			chainIDs, err := parseChainIDs(chainIDsStr)
			if err != nil {
				return fmt.Errorf("failed to parse chain IDs: %v", err)
			}

			// 解析 value
			val := new(big.Int)
			_, success := val.SetString(value, 10)
			if !success {
				return fmt.Errorf("invalid value: %s", value)
			}

			// 构造参数
			crossSendArgs := []scripts.CrossSendArg{
				{
					ChainIDs: chainIDs,
					Value:    val,
				},
			}

			ctx := context.Background()
			if err := scripts.CrossSend(ctx, cfg, crossSendArgs); err != nil {
				return fmt.Errorf("cross-send failed: %v", err)
			}

			fmt.Println("CrossSend completed successfully")
			return nil
		},
	}

	// 添加 cross-send 命令的标志
	crossSendCmd.Flags().StringVar(&chainIDsStr, "chain-ids", "", "Comma-separated chain IDs (e.g., '1,2,3')")
	crossSendCmd.Flags().StringVar(&value, "value", "0", "Value to send")
	crossSendCmd.MarkFlagRequired("chain-ids")
	crossSendCmd.MarkFlagRequired("value")

	// RegistEICN 命令
	registEICNCmd := &cobra.Command{
		Use:   "regist-eICN",
		Short: "Regist EICN",
		RunE: func(cmd *cobra.Command, args []string) error {
			sourceConfig, err := config.LoadConfig(cfgFile)
			if err != nil {
				return fmt.Errorf("failed to load config: %v", err)
			}
			targetConfig, err := config.LoadConfig(targetConfigFile)
			if err != nil {
				return fmt.Errorf("failed to load config: %v", err)
			}
			if err := scripts.RegistEICN(sourceConfig, targetConfig); err != nil {
				return fmt.Errorf("regist-eICN failed: %v", err)
			}
			return nil
		},
	}
	registEICNCmd.Flags().StringVar(&targetConfigFile, "target-config", "config.yaml", "Target config file path")
	registEICNCmd.MarkFlagRequired("target-config")
	// 添加子命令到根命令
	rootCmd.AddCommand(
		deployCmd,
		deployLibrariesCmd,
		crossSendCmd,
		registEICNCmd,
	)

	// 执行
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
