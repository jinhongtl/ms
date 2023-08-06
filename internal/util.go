package internal

import (
	"fmt"

	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitConfig() {
	//配置文件名称
	viper.SetConfigName("config")
	//配置文件类型
	viper.SetConfigType("yaml")
	//读取配置文件的范围
	viper.AddConfigPath("/etc/appname/")
	viper.AddConfigPath("$HOME/.appname")
	viper.AddConfigPath(".")
	viper.AddConfigPath(".")
	//开始读取
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	//将配置文件设置到结构体中
	viper.Unmarshal(&AppConfig)
}

func InitDb() {
	dsn := "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	dbConf := AppConfig.Database
	dsn = fmt.Sprintf(dsn, dbConf.User, dbConf.Password, dbConf.Host, dbConf.Port, dbConf.DbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("could not connect mysql db")
	}
	Db = db
}

func InitLogger() {
	Logger = logrus.New()
	file, err := os.Create(AppConfig.Log.Path)
	if err != nil {
		panic("could not open log file in path " + AppConfig.Log.Path)
	}
	Logger.SetOutput(file)
	Logger.SetLevel(logrus.Level(AppConfig.Log.Level))
}
