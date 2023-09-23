package controller

import (
	_ "github.com/LittleMikle/sber_it/docs"
	"github.com/LittleMikle/sber_it/pkg/service"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Controller struct {
	services *service.Service
}

func NewController(services *service.Service) *Controller {
	return &Controller{
		services: services,
	}
}

func (h *Controller) InitRoutes() *gin.Engine {
	router := gin.New()
	lists := router.Group("/lists")
	{
		lists.POST("/create", h.createList)
		lists.GET("/lists", h.getLists)
		lists.PUT("/:id", h.updateList)
		lists.DELETE("/:id", h.deleteList)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router
}
