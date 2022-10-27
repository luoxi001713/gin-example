package config

import (
	"fmt"
	"encoding/json"
	"gorm.io/gorm"
	"os"
	"path/filepath"
)

var DB *gorm.DB
var Conf DBConfig
var AllConf Config

func Init(){

	path, _ := filepath.Abs("./config/config.json")
	file, err := os.Open(path)
	if err != nil {
		os.Exit(0)
	}
	decoder := json.NewDecoder(file)
	AllConf = Config{}
	err = decoder.Decode(&AllConf)
	if err != nil {
		os.Exit(0)
	}
	Conf = AllConf.DBConfig
	return
}

//Config include all config info
type Config struct {
	DBConfig     DBConfig          `json:"db"`
}

// DBConfig represents db configuration
type DBConfig struct {
	Host     string `json:"db_host"`
	Port     int    `json:"db_port"`
	User     string `json:"db_user"`
	DBName   string `json:"db_name"`
	Password string `json:"db_password"`
}

func DbURL() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		Conf.User,
		Conf.Password,
		Conf.Host,
		Conf.Port,
		Conf.DBName,
	)
}
