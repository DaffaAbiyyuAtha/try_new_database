package routers

import (
	"fazztrack/backend/controllers"

	"github.com/gin-gonic/gin"
)

func UserRouter(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/", controllers.ListAllUser)
	routerGroup.GET("/:id", controllers.DetailUser)
	routerGroup.POST("/", controllers.CreateUser)
	routerGroup.PATCH("/:id", controllers.UpdateUser)
	routerGroup.DELETE("/:id", controllers.DeleteUser)
}