package socket

import (
	"cloud-collector/config"
	"cloud-collector/define"
	"cloud-collector/logger"
	"context"
	"github.com/TencentBlueKing/bkmonitor-datalink/pkg/libgse/gse"
	"os"
	"sync"
)

var (
	defaultWorkerNum   = 1
	defaultGseClient   *gse.GseSimpleClient
	defaultQueueBuffer = 500
	Name               = "GSEServer"
	GlobalMsgCh        chan gse.GseMsg
)

type GseSocket struct {
	c   *config.Socket
	ch  chan gse.GseMsg
	wg  *sync.WaitGroup
	ctx context.Context
	//cancel context.CancelFunc
}

func NewGseSocketServer(c *config.Socket) *GseSocket {
	defaultGseClient = gse.NewGseSimpleClient()
	socketPath := c.SocketPath
	if socketPath != "" {
		defaultGseClient.SetAgentHost(socketPath)
	} else {
		defaultGseClient.SetAgentHost(define.DefaultSocketPath)
		logger.Warn("use default socket path:", define.DefaultSocketPath)
	}
	// 根据配置进行消息队列缓冲区的设置
	if c.QueueBuffer != 0 {
		GlobalMsgCh = make(chan gse.GseMsg, c.QueueBuffer)
	} else {
		GlobalMsgCh = make(chan gse.GseMsg, defaultQueueBuffer)
	}

	return &GseSocket{
		c:  c,
		ch: GlobalMsgCh,
		wg: &sync.WaitGroup{},
	}
}

func (g *GseSocket) StartServer(c context.Context) {
	g.ctx = c
	g.wg.Add(1)
	go func() {
		defer g.wg.Done()
		g.startGseServer()
	}()

	<-g.ctx.Done()
	g.wg.Wait()

}

// startGseServer 启动GSE服务，并且监听 MsgQueue 等待数据上报
func (g *GseSocket) startGseServer() {
	err := defaultGseClient.Start()
	if err != nil {
		logger.Errorf("start gse server error: %v\n", err)
		os.Exit(1)
	}
	logger.Debugln("start gse server success!")

	if g.c.Worker != 0 {
		defaultWorkerNum = g.c.Worker
	}

	for i := 0; i < defaultWorkerNum; i++ {
		g.wg.Add(1)
		go g.monitorMsg(i)
	}
}

func (g *GseSocket) monitorMsg(workerId int) {
	for {
		select {
		case msg := <-g.ch:
			if err := defaultGseClient.Send(msg); err != nil {
				logger.Errorf("send msg error: %v\n", err)
			}
			logger.Debugf("send msg success: %v\n", msg)
		// 监听父进程（主进程的中断信号）
		case <-g.ctx.Done():
			g.wg.Done()
			logger.Infof("gse goroutine:%d exit", workerId)
			return
		}
	}
}
