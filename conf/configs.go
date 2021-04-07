package conf

//
//import (
//	"errors"
//	"github.com/BurntSushi/toml"
//	"log"
//	"os"
//)
//
//const (
//	configFileName = "application.toml"
//)
//
//var (
//	confPath string = "/Users/yuzhang/go/src/examination-sys/conf/application.toml"
//	Conf     = &Config{}
//)
//
//
//type Config struct {
//	Database    Database         `toml:"database"`
//}
//
//type Database struct {
//	Host     string `toml:"host"`
//	Port     int    `toml:"port"`
//	DBName   string `toml:"dbname"`
//	User     string `toml:"user"`
//	Password string `toml:"password"`
//}
//
//
//
//
//// ******************* 载入 config *********************
//
//func init()  {
//	// 配置初始化
//	local()
//}
//
//// 本地配置
//func local() error {
//	if _, err := toml.DecodeFile(confPath, &Conf); err != nil {
//		return err
//	}
//	return nil
//}
//
//
//func load() (err error) {
//	var (
//		tmpConf Config
//	)
//	if _, err := os.Stat(confPath); os.IsNotExist(err) {
//		return errors.New("Config file does not exist.")
//	} else if err != nil {
//		return  err
//	}
//	//log.Info("%+v", s)
//	if _, err := toml.Decode(confPath, &tmpConf); err != nil {
//		log.Println(err.Error())
//		return errors.New("could not decode config")
//	}
//
//	Conf = &tmpConf
//
//	return
//}
