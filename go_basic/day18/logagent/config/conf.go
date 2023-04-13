package config

type Config struct {
	KafkaConf `ini:"Kafka"`
	EtcdConf  `ini:"Etcd"`
	EsConf    `ini:"es"`
}
type KafkaConf struct {
	Addr string `ini:"Addr"`
}

type EtcdConf struct {
	Addr    string `ini:"Address"`
	Key     string `ini:"Collect_logEntry_key"`
	Timeout int    `ini:"Timeout"`
}
type EsConf struct {
	Addr       string `ini:"addr"`
	ChanMax    int    `ini:"chan_max"`
	RoutineMax int    `ini:"msgToEs_routine_max"`
}
