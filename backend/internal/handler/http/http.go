package http

import (
	"context"
	"doremi/internal/model"
	"doremi/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/hellofresh/health-go/v5"
	"net/http"
)

type Http struct {
	service service.Service
}

func NewHttpHandler(service service.Service) *Http {
	return &Http{
		service: service,
	}
}

func (h *Http) CheckHealth() http.HandlerFunc {
	healthController, _ := health.New(health.WithComponent(health.Component{
		Name:    "mongodb",
		Version: "v1.0",
	}), health.WithChecks(health.Config{
		Name: "mongodb",
		Check: func(ctx context.Context) error {
			return nil
		},
	},
	))

	return healthController.HandlerFunc
}

func (h *Http) Register(c *gin.Context) {
	var user model.User
	if err := c.Bind(&user); err != nil {
		ErrorHandler(c, err, http.StatusInternalServerError)
		return
	}

	if err := h.service.User.CreateNewUser(user, c); err != nil {
		ErrorHandler(c, err, http.StatusInternalServerError)
		return
	}

	SuccessHandler(c, model.SuccessUserCreation)
}

func (h *Http) Login(c *gin.Context) {
	return
}
