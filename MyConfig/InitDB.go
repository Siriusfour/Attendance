package MyConfig

import (
	"GrpcTest/Model"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

func InitDB() (*gorm.DB, error) {

	//查看当前环境级别
	LogMode := logger.Info
	if !viper.GetBool("Model.develop") {
		LogMode = logger.Error
	}
	//创建一个gorm实例
	db, err := gorm.Open(mysql.Open(viper.GetString("DB.DNS")), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			TablePrefix:   "sys_",
		},
		Logger: logger.Default.LogMode(LogMode),
	})

	if err != nil {
		return nil, err
	}

	SqlDB, _ := db.DB()
	SqlDB.SetMaxIdleConns(viper.GetInt("DB.MaxIdleConns"))
	SqlDB.SetMaxOpenConns(viper.GetInt("DB.MaxOpenConns"))
	SqlDB.SetConnMaxLifetime(time.Hour)

	//将数据库的表与user对象同步，有则修改，没有创建
	err = db.AutoMigrate(&Model.User{}, &Model.Application{})
	if err != nil {
		panic(err)
	}

	return db, nil
}
