package server

import (
	"doremi/internal/config"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func StartServer(cfg config.Config, g *gin.Engine) error {
	srv := &http.Server{
		Addr:         cfg.Host + ":" + cfg.Port,
		Handler:      g,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	err := srv.ListenAndServe()

	return err
}
