package handler

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) Routes() *gin.Engine {
	r := gin.New()

	v1 := r.Group("/v1")
	{
		v1.GET("/health", func(c *gin.Context) {
			h.Http.CheckHealth()
		})
		authGroup := v1.Group("/auth")
		{
			authGroup.POST("/register", h.Http.Register)
			authGroup.POST("/login", h.Http.Login)
		}
	}

	return r
}
