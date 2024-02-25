package handler

import (
	"context"
	"github.com/hellofresh/health-go/v5"
	"net/http"
)

func CheckHealth() http.HandlerFunc {
	h, _ := health.New(health.WithComponent(health.Component{
		Name:    "mongodb",
		Version: "v1.0",
	}), health.WithChecks(health.Config{
		Name: "mongodb",
		Check: func(ctx context.Context) error {

			return nil
		},
	},
	))

	return h.HandlerFunc
}
