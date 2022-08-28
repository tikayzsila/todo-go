package handler

import (
	"todo-app/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {

	return &Handler{services: services}

}

func (h *Handler) InitRoutes() *gin.Engine {

	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/login", h.logIn)
		auth.POST("/sign-up", h.signUp)
	}

	api := router.Group("/api", h.userIdentity)
	{
		lists := api.Group("/lists")
		{
			lists.GET("/", h.getAllLists)
			lists.POST("/", h.createList)
			lists.GET("/:id", h.gitListById)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)
			items := lists.Group(":id/items")
			{
				items.GET("/", h.getAllItem)
				items.POST("/", h.createItem)
				items.GET("/:id", h.gitItemById)
				items.PUT("/:id", h.updateItem)
				items.DELETE("/:id", h.deleteItem)
			}
		}

	}

	return router

}
