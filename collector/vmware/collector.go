package vmware

import (
	"cloud-collector/common"
	"cloud-collector/logger"
	"context"
	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25/mo"
	"net/url"
	"sync"
)

// CollectorTask 实际的任务执行体，可以理解为一个 CollectorTask 就是一个 VMCloud
type CollectorTask struct {
	conf *Cloud
	ctx  context.Context

	wg     sync.WaitGroup
	client *govmomi.Client
}

func NewCollectorTask(c *Cloud, ctx context.Context) *CollectorTask {
	return &CollectorTask{
		conf: c,
		ctx:  ctx,
		wg:   sync.WaitGroup{},
	}
}

func (c *CollectorTask) getClient() *govmomi.Client {
	if c.client != nil {
		return c.client
	}
	password := common.DecryptPassword(c.conf.Password, aesKey)
	u := &url.URL{
		Scheme: "https",
		Host:   c.conf.Server,
		Path:   "/sdk",
	}
	u.User = url.UserPassword(c.conf.Account, password)
	client, err := govmomi.NewClient(c.ctx, u, true)
	if err != nil {
		logger.Errorf("unable to create vmware client: %v\n", err)
		return nil
	}
	c.client = client
	return client
}

// collectCluster 收集集群的情况
func (c *CollectorTask) collectCluster() {

}

// CollectHost 收集实体的宿主机情况
func (c *CollectorTask) CollectHost() {
	client := c.getClient()
	if client == nil {
		logger.Errorf("client is nil\n")
		return
	}

	kind := []string{"HostSystem"}
	manager := view.NewManager(client.Client)
	v, err := manager.CreateContainerView(c.ctx, client.ServiceContent.RootFolder, kind, true)
	if err != nil {
		logger.Errorf("unable to create container view: %v\n", err)
		return
	}
	// 退出前关闭 ContainerView
	defer func() {
		if err := v.Destroy(c.ctx); err != nil {
			logger.Errorf("unable to destroy container view: %v\n", err)
			return
		}
	}()

	var hosts []mo.HostSystem
	if err := v.Retrieve(c.ctx, kind, []string{"summary", "datastore"}, &hosts); err != nil {
		logger.Errorf("unable to retrieve hosts, error:%s\n", err)
		return
	}
	// 判定一下 hosts 是否为空的情况
	if len(hosts) == 0 {
		logger.Warnln("no hosts found, nothing to collect")
		return
	}

	//perfManager := performance.NewManager(client.Client)
	//counters, err := perfManager.CounterInfoByName(c.ctx)
	//if err != nil {
	//	logger.Errorf("unable to retrieve perf counters, error:%s\n", err)
	//	return
	//}

}

// CollectVM 收集 虚拟机的情况
func (c *CollectorTask) CollectVM() {}

// CollectStorage 收集 磁盘
func (c *CollectorTask) CollectStorage() {}

// Start 启动采集任务
func (c *CollectorTask) Start() {}
