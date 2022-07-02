package config

type AppConfig struct {
	KafkaConfig `ini:"kafka"`
	Taillog `ini:"taillog"`
}

type KafkaConfig struct {
	Address string `ini:"address"`
	Topic string `ini:"topic"`
}

type Taillog struct {
	FieName string `ini:"fiename"`
}
type EtcdConfig struct {
	Address string `ini:"address"`
	TimeOut int `ini:"timeout"`
}