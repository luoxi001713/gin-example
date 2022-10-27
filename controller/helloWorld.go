package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code            int                  `json:"code"`
	Success         bool                 `json:"success"`
	Message         string               `json:"message"`
	Data            interface{}          `json:"data"`
}

type HelloWorldRequest struct {
	Name         string       `json:"name" form:"name" binding:"required"`
}

type HelloWorldResponse struct {
	Greeting         string       `json:"greeting"`
}


func HelloWorld(c *gin.Context) {

	var req HelloWorldRequest
    var resp Response

	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"msg":err.Error(),
		})
		c.Abort()
		return
	}

	var data HelloWorldResponse
	data.Greeting = fmt.Sprintf("Hello %s", req.Name)

	resp.Code = http.StatusOK
	resp.Data = data
	resp.Success = true
	c.JSON(resp.Code, resp)
}