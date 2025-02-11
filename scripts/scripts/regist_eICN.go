package scripts

import (
	"github.com/kimroniny/SuperRunner-eICN-eth2/client"
	"github.com/kimroniny/SuperRunner-eICN-eth2/config"
)

func RegistEICN(sourceConfig *config.Config, targetConfig *config.Config) error {
	transmitterClient := client.NewTransmitterClient(nil)
	sourceTransmitterURL := sourceConfig.HTTP.BaseURL
	targetTransmitterURL := targetConfig.HTTP.BaseURL
	sourceChainID := sourceConfig.Chain.ID

	err := transmitterClient.RegisterEICN(sourceTransmitterURL, sourceChainID, targetTransmitterURL)
	if err != nil {
		return err
	}

	return nil
}
