package main

import (
	"fmt"

	"github.com/fsnotify/fsnotify"

	"github.com/spf13/viper"
)

type App struct { //注意：一定要配置一个全局的，否则无法解析到结构体
	Mysql `mapstructure:"mysql"` //指针或者值嵌入均可
}
type Mysql struct {
	Dbname   string `mapstructure:"dbname"`
	Port     int64  `mapstructure:"port"`
	User     string `mapstructure:"dbuser"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"dbhost"`
}

func main() {

	viper.SetConfigName("config") //配置文件名称 对于本地文件不必带后缀
	viper.AddConfigPath("../")    //可以添加多个配置文件路径，依次寻找，找到第一个后将不再寻找
	viper.AddConfigPath(".")      // 添加配置文件路径
	err := viper.ReadInConfig()   // 读取配置文件
	if err != nil {
		fmt.Println(err)
		return
	}
	app := new(App)
	err = viper.Unmarshal(app)
	if err != nil {
		fmt.Println(err)
		return
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config changes!", in.Name)
	})
	fmt.Println(app.Port)
	select {}

}
