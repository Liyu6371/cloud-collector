package config

import (
	"cloud-collector/collector/vmware"
	"cloud-collector/collector/winstack"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

var globalConfig *Config

type Config struct {
	Logger           *Logger           `yaml:"logger"`
	Socket           *Socket           `yaml:"socket"`
	CloudCollectTask *CloudCollectTask `yaml:"cloud_collect_task"`
}

type CloudCollectTask struct {
	WinStackTask *winstack.WinStack  `yaml:"win_stack"`
	VMWare       *vmware.VMCollector `yaml:"vm_ware"`
	//VMWareTask   *vmware.VMWareTask `yaml:"vmware_task"`
}

func ParseConfig(p string) (*Config, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("cannot get working directory: %v", err)
	}
	confPath := filepath.Join(dir, p)
	content, err := os.ReadFile(confPath)
	if err != nil {
		return nil, fmt.Errorf("cannot read config file: %s, error: %v", confPath, err)
	}
	err = yaml.Unmarshal(content, &globalConfig)
	if err != nil {
		return nil, fmt.Errorf("yaml unmarshal error: %v", err)
	}
	return globalConfig, nil
}

// GetConfig 获取全局配置文件
func GetConfig() *Config {
	return globalConfig
}
