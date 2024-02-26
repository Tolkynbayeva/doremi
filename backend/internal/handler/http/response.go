package http

import (
	"doremi/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Status      string `json:"status"`
	Description string `json:"description"`
}

func ErrorHandler(c *gin.Context, err error, code int) {
	response := Response{
		Status:      model.ErrorMsg,
		Description: err.Error(),
	}

	c.JSON(code, response)
}

func SuccessHandler(c *gin.Context, operation string) {
	response := Response{
		Status:      model.SuccessMsg,
		Description: operation,
	}

	c.JSON(http.StatusOK, response)
}
