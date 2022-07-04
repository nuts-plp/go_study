package config

type AppConfig struct {
	KafkaConfig   `ini:"kafka"`
	TaillogConfig `ini:"taillog"`
	EtcdConfig    `ini:"etcd"`
}

type KafkaConfig struct {
	Address     string `ini:"address"`
	Topic       string `ini:"topic"`
	ChanMaxSize int    `ini:"chan_max_size"`
}

type TaillogConfig struct {
	FieName string `ini:"filename"`
}
type EtcdConfig struct {
	Address string `ini:"address"`
	TimeOut int    `ini:"timeout"`
	Key     string `ini:"collect_log_config"`
}
