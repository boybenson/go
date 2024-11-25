package routes

import (
	"example/go/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine){
	router.GET("/poems", controllers.HandleGetPoems)

	router.POST("/poem", controllers.HandleCreatePoem)

	router.GET("/poems/:id", controllers.HandleGetPoem)

	router.DELETE("/poems/:id",  controllers.HandleDeletePoem)
}