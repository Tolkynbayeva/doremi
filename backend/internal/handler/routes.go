package handler

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (h *Handler) Routes() *gin.Engine {
	r := gin.New()

	swaggerURL := ginSwagger.URL("http://localhost:8181/swagger/doc.json")

	v1 := r.Group("/v1")
	{
		v1.GET("/health", func(c *gin.Context) {
			h.Http.CheckHealth()
		})
		v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, swaggerURL))
		authGroup := v1.Group("/auth")
		{
			authGroup.POST("/register", h.Http.Register)
			authGroup.POST("/login", h.Http.Login)
		}
	}

	return r
}
