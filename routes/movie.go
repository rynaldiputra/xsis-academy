package routes

import (
	"theater/handlers"

	"github.com/gin-gonic/gin"
)

func MovieRoutes(api *gin.RouterGroup, handler *handlers.MovieHandler) {
	api.GET("/", handler.Get)
	api.GET("/:id", handler.Find)
	api.POST("/", handler.Create)
	api.PUT("/:id", handler.Edit)
	api.DELETE("/:id", handler.Delete)
}
