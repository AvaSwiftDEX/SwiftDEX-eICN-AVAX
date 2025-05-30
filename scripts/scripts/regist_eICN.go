package scripts

import (
	"strconv"

	"github.com/kimroniny/SuperRunner-eICN-eth2/client"
	"github.com/kimroniny/SuperRunner-eICN-eth2/config"
)

func RegistEICN(sourceConfig *config.Config, targetConfig *config.Config) error {
	transmitterClient := client.NewTransmitterClient(nil)
	sourceTransmitterURL := "http://" + sourceConfig.HTTP.Host + ":" + strconv.Itoa(int(sourceConfig.HTTP.Port))
	targetTransmitterURL := "http://" + targetConfig.HTTP.Host + ":" + strconv.Itoa(int(targetConfig.HTTP.Port))
	sourceChainID := sourceConfig.Chain.ID

	err := transmitterClient.RegisterEICN(sourceTransmitterURL, sourceChainID, targetTransmitterURL)
	if err != nil {
		return err
	}

	return nil
}
