package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Init() (err error) {
	//viper.SetConfigFile("filename.type")指定配置文件
	viper.SetConfigName("config") //指定配置文件名称，不需带后缀
	viper.SetConfigType("yaml")   //专用于远程获取配置信息时指定配置文件类型
	viper.AddConfigPath("./settings/")
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Printf("viper.ReadInConfig failed,err:%v \n", err)
		return err
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config changed")
	})
	return
}
