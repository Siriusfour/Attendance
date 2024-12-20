package InitConfig

import (
	"fmt"
	"github.com/spf13/viper"
)

func InitConfig() {
	viper.SetConfigName("setting")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("D:/go/GOPATH/grpc-test/InitConfig")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("加载配置文件出错：%s", err))
	}

	//初始化数据库
	DB, err = InitDB()
	if err != nil {
		panic("DB InitConfig fail:" + err.Error())
	}

	//初始化日志实例
	Logger = InitLogger()

}
