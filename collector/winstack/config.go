// WinStackCollect, 暂时不考虑有多个 WinStackCloud 实例的情况；暂时只考虑一个的状态下
// 周期调度上报数据到 gse

package winstack

// WinStack 考虑到采集任务运行时间可能比较长，在这里仅支持50并发量的 WinStack 任务运行
// 如后续机器性能不足的情况下需要调小 Concurrency 数值
type WinStack struct {
	Concurrency int      `yaml:"concurrency"`
	Clouds      *[]Cloud `yaml:"clouds"`
}

type Cloud struct {
	ID       int      `yaml:"id"`
	Period   string   `yaml:"period"`
	Timeout  string   `yaml:"timeout"`
	Account  string   `yaml:"account"`
	Password string   `yaml:"password"`
	Server   string   `yaml:"server"`
	Cluster  *Cluster `yaml:"cluster"`
	Host     *Host    `yaml:"host"`
	Storage  *Storage `yaml:"storage"`
	VM       *VM      `yaml:"vm"`
}

type Cluster struct {
	DataId    int32    `yaml:"data_id"`
	Instances []string `yaml:"instances"`
}

type Host struct {
	DataId    int32    `yaml:"data_id"`
	Instances []string `yaml:"instances"`
}

type Storage struct {
	DataId    int32    `yaml:"data_id"`
	Instances []string `yaml:"instances"`
}

type VM struct {
	DataId    int32    `yaml:"data_id"`
	Instances []string `yaml:"instances"`
}
