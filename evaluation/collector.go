package evaluation

import (
	"time"

	"github.com/kimroniny/SuperRunner-eICN-eth2/SR2PC"
	"github.com/kimroniny/SuperRunner-eICN-eth2/config"
)

type Collector struct {
	Config *config.Config
}

func NewCollector(config *config.Config) *Collector {
	return &Collector{Config: config}
}

func (c *Collector) CollectCM(cm *SR2PC.SR2PCCrossMessage, status uint8, timestamp time.Duration) {

}

func (c *Collector) TSIssue(transactionHash string, status uint8) {

}
