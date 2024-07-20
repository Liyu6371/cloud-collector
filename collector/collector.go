package collector

import (
	"cloud-collector/collector/vmware"
	"cloud-collector/collector/winstack"
	"cloud-collector/config"
	"cloud-collector/logger"
	"context"
	"sync"
)

var Name = "CollectorServer"

type CloudCollector interface {
	Run(ctx context.Context)
}

type Collector struct {
	c   *config.CloudCollectTask
	wg  *sync.WaitGroup
	ctx context.Context
}

func NewCollectorServer(c *config.CloudCollectTask) *Collector {
	return &Collector{
		c:  c,
		wg: &sync.WaitGroup{},
	}
}

func (c *Collector) StartServer(ctx context.Context) {
	c.ctx = ctx
	// 启动 VM 处理
	c.runCloudCollector(c.c.VMWare, vmware.Name)
	// 启动 WinStack 处理
	c.runCloudCollector(c.c.WinStackTask, winstack.Name)
	<-c.ctx.Done()
	c.wg.Wait()
	logger.Info("CollectorServer stopped")
}

func (c *Collector) runCloudCollector(collector CloudCollector, name string) {
	if collector == nil {
		logger.Errorf("collector is nil,  server name is: %s\n", name)
		return
	}
	c.wg.Add(1)
	go func(collector CloudCollector) {
		defer c.wg.Done()
		collector.Run(c.ctx)
	}(collector)
	logger.Info("collector :%s is running", name)
}
