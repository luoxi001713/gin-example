package main

import (

	"github.com/luoxi001713/gin-example/config"
	"github.com/luoxi001713/gin-example/router"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {

	config.Init()

	config.DB, _ = gorm.Open(mysql.Open(config.DbURL()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	r := router.SetupRouter()
	r.Run(":9090")
}
