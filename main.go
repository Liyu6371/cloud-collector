package main

import (
	"cloud-collector/config"
	"cloud-collector/define"
	"cloud-collector/logger"
	"flag"
	"fmt"
	"os"
)

var (
	fileFlag    = flag.String("c", "", "relative path of configuration file ")
	versionFlag = flag.Bool("v", false, "print version and exit")
	testFlag    = flag.Bool("test", false, "test mode")
)

func main() {
	flag.Parse()
	// 查看版本
	if *versionFlag {
		fmt.Println(define.Version)
		os.Exit(0)
	}
	// 解析配置文件
	if *fileFlag == "" {
		fmt.Println("config file should not be empty")
		fmt.Println("Usage: ./cloud-collector -c /path/to/configuration.yml")
		os.Exit(1)
	}
	c, err := config.ParseConfig(*fileFlag)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// 初始化 logger
	logger.SetLogByConfig(c.Logger)
	// cloudCollectorService 初始化
	svc := NewCloudCollectorService(c)
	// 拉起所有服务
	svc.Run()
}
