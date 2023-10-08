package routers

import (
	"github.com/beruang43221/assignment3-015/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.PUT("/update/:id", controller.UpdateWeather)

	return r

}