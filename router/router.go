package router

import (
	"mojito/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	
	v1 := r.Group("/v1")
	{
		// v1.GET("todo", controller.GetAllLost)
		v1.POST("lost", controller.CreateLost)
		v1.GET("lost/:id", controller.GetLost)
		v1.PUT("lost/:id", controller.UpdateLost)
		v1.DELETE("lost/:id", controller.DeleteLost)
	}

	return r
}