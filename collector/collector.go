package collector

import (
	"cloud-collector/config"
	"context"
)

var Name = "CollectorServer"

type Collector struct {
	c *config.CloudCollectTask
}

func NewCollectorServer(c *config.CloudCollectTask) *Collector {
	return &Collector{
		c: c,
	}
}

func (c *Collector) StartServer(ctx context.Context) {}
