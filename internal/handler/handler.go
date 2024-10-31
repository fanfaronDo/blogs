package handler

import (
	"github.com/fanfaronDo/blogs/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(h.middlewareCORS())
	posts := router.Group("/posts")
	{
		posts.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "pong"})
			return
		})

		posts.GET("", h.getPosts)
		posts.GET("/:id", h.getPost)

	}

	admin := router.Group("/admin")
	{
		admin.POST("", h.createPost)
		admin.PATCH("", h.updateProst)
	}

	return router
}
