package http

import (
	"context"
	"doremi/internal/model"
	"doremi/internal/service"
	"errors"
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

// @Summary Register a new user
// @Description Register a new user with username, email, and password
// @ID register-user
// @Accept  json
// @Produce  json
// @Param   user  body    model.User  true  "User Registration Data"
// @Success 200  {object}  model.Response  "User successfully registered"
// @Failure 500  {object}  model.Response  "Internal Server Error"
// @Router /v1/auth/register [post]
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
	var creds model.User
	if err := c.Bind(&creds); err != nil {
		ErrorHandler(c, err, http.StatusInternalServerError)
		return
	}

	user, err := h.service.User.CheckUserCredentials(creds, c)
	if err != nil {
		if errors.Is(err, model.ErrIncorrectPassword) {
			ErrorHandler(c, err, http.StatusBadRequest)
		} else {
			ErrorHandler(c, err, http.StatusInternalServerError)
		}
		return
	}

	token, err := h.service.User.JwtAuth(user)
	if err != nil {
		ErrorHandler(c, err, http.StatusInternalServerError)
		return
	}

	c.SetCookie("Authorization", token, 3600*24*30, "", "", false, true)
	SuccessHandler(c, model.SuccessLogin)
}
