package conf

type AppConf struct {
	KafkaConf `ini:"kafka"`
	Taillog   `ini:"taillog"`
	EtcdConf  `ini:"etcd"`
}
type KafkaConf struct {
	Address     string `ini:"address"`
	Topic       string `ini:"topic"`
	ChanMaxSize int    `ini:"chan_max_size"`
}
type Taillog struct {
	Filename string `ini:"path"`
}
type EtcdConf struct {
	Address string `ini:"addr"`
	Timeout int    `ini:"timeout"`
	Key     string `ini:"collect_log_key"`
}
