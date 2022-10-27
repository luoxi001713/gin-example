package router

import (

	"github.com/gin-gonic/gin"
	"github.com/luoxi001713/gin-example/controller"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()

	api := r.Group("platform/api")
	{
		api.GET("helloWorld",controller.HelloWorld)
	}
    return r	
}