package main

import (
	"cloud-collector/collector"
	"cloud-collector/config"
	"cloud-collector/logger"
	"cloud-collector/socket"
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type Server interface {
	StartServer(ctx context.Context)
}

type CloudCollectorService struct {
	wg              *sync.WaitGroup
	ctx             context.Context
	cancel          context.CancelFunc
	socketServer    *socket.GseSocket
	collectorServer *collector.Collector
}

func NewCloudCollectorService(c *config.Config) *CloudCollectorService {
	cloudCollectorServer := &CloudCollectorService{
		wg: &sync.WaitGroup{},
	}
	ctx, cancel := context.WithCancel(context.Background())
	cloudCollectorServer.ctx = ctx
	cloudCollectorServer.cancel = cancel
	if c.Socket != nil {
		cloudCollectorServer.socketServer = socket.NewGseSocketServer(c.Socket)
	}
	if c.CloudCollectTask != nil {
		cloudCollectorServer.collectorServer = collector.NewCollectorServer(c.CloudCollectTask)
	}
	return cloudCollectorServer
}

func (svc *CloudCollectorService) Run() {
	// 非测试模式下启动 GSE Client
	if !*testFlag {
		svc.runServer(svc.socketServer, socket.Name)
	}
	// 启动云监控采集服务
	svc.runServer(svc.collectorServer, collector.Name)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	// 先通过上下文ctx通知所有的子进程退出
	svc.cancel()
	// 等待所有的子进程退出后再退出主进程
	svc.wg.Wait()
}

func (svc *CloudCollectorService) runServer(server Server, name string) {
	if server == nil {
		logger.Errorf("server is nil, process exit, server name is: %s\n", name)
		os.Exit(1)
	}
	svc.wg.Add(1)
	go func(s Server) {
		defer svc.wg.Done()
		s.StartServer(svc.ctx)
	}(server)
	logger.Info("server %s is running", name)
}
