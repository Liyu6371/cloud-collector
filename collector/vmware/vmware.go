package vmware

import (
	"cloud-collector/logger"
	"context"
	"sync"
	"time"
)

var (
	wg     sync.WaitGroup    // goroutine 控制
	period = time.Minute * 5 // 默认运行周期
)

func (v *VMCollector) Run(ctx context.Context) {
	defer logger.Infoln("VMCollector shutting down")
	if len(*v.Clouds) == 0 {
		logger.Infoln("VMCollector Run: No Cloud found, Nothing to do")
		return
	}
	// VMWare_Collector 采集为周期 Ticker 调度的任务，可能存在资源消耗比较大的情况
	// 一个 VMWare 实例里面可能存在 N 个需要采集的目标，因此需要对进行采集的 VMWare 实例数量进行限制
	// 当实例的数量大于并发限制数量的时候输出告警信息
	if len(*v.Clouds) > v.Concurrency {
		logger.Warnln("VMCollector Run: Number of Cloud Foundry is greater than Concurrency")
		logger.Warnln("VMCollector Run: The rest of Cloud Foundry will be ignored")
	}

	// 拉起任务
	count := 0
	for _, cloud := range *v.Clouds {
		count += 1
		if count > v.Concurrency {
			logger.Warnln("VMCollector Run: Too many cloud found")
			break
		} else {
			wg.Add(1)
			go process(cloud, ctx)
		}
	}

	<-ctx.Done()
	logger.Infoln("VMCollector Run: catch ctx.Done signal, waiting for all goroutines to finish")
	wg.Wait()
	logger.Infoln("VMCollector Run: All goroutines finished")
}

func process(c Cloud, ctx context.Context) {
	// 退出前必须要处理 信号量 以及 wg
	//defer func() {
	//	wg.Done()
	//}()
	//
	//ticker := time.NewTicker(time.Minute)
	//defer ticker.Stop()
	//
	//for {
	//	select {
	//	case <-ctx.Done():
	//		return
	//	case <-ticker.C:
	//		fmt.Printf("processing %+v", c)
	//	}
	//}
}
