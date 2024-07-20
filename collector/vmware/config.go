package vmware

var Name = "VMWareCollector"

type VMCollector struct {
	Concurrency int      `yaml:"concurrency"`
	Clouds      *[]Cloud `yaml:"clouds"`
}

type Cloud struct {
	Id       int      `yaml:"id"`
	Server   string   `yaml:"server"`
	Account  string   `yaml:"account"`
	Password string   `yaml:"password"`
	Cluster  *Cluster `yaml:"cluster"`
	Host     *Host    `yaml:"host"`
	Storage  *Storage `yaml:"storage"`
	VM       *VM      `yaml:"vm"`
}

type Cluster struct {
	ClusterMetricNamespace string     `yaml:"cluster_metric_namespace"`
	ClusterMetricDataId    int32      `yaml:"cluster_metric_data_id"`
	ClusterEventDataId     int32      `yaml:"cluster_event_data_id"`
	ClusterInstances       *[]string  `yaml:"cluster_instances"`
	ClusterMetrics         *[]Metrics `yaml:"cluster_metrics"`
}

type Host struct {
	HostMetricNamespace string     `yaml:"host_metric_namespace"`
	HostMetricDataId    int32      `yaml:"host_metric_data_id"`
	HostEventDataId     int32      `yaml:"host_event_data_id"`
	HostInstances       *[]string  `yaml:"host_instances"`
	HostMetrics         *[]Metrics `yaml:"host_metrics"`
}

type Storage struct {
	StorageMetricNamespace string     `yaml:"storage_metric_namespace"`
	StorageMetricDataId    int32      `yaml:"storage_metric_data_id"`
	StorageEventDataId     int32      `yaml:"storage_event_data_id"`
	StorageInstances       *[]string  `yaml:"storage_instances"`
	StorageMetrics         *[]Metrics `yaml:"storage_metrics"`
}

type VM struct {
	VMMetricNamespace string     `yaml:"vm_metric_namespace"`
	VMMetricDataId    int32      `yaml:"vm_metric_data_id"`
	VMEventDataId     int32      `yaml:"vm_event_data_id"`
	VMInstances       *[]string  `yaml:"vm_instances"`
	VMMetrics         *[]Metrics `yaml:"vm_metrics"`
}

type Metrics struct {
	Alias  string `yaml:"alias"`
	Metric string `yaml:"metric"`
}
