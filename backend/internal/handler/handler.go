package handler

import (
	"doremi/internal/handler/grpc"
	"doremi/internal/handler/http"
	"doremi/internal/service"
)

type Handler struct {
	Http    *http.Http
	Grpc    *grpc.Grpc
	service service.Service
}

func NewHandler(service service.Service) *Handler {
	return &Handler{
		service: service,
		Http:    http.NewHttpHandler(service.User),
		Grpc:    grpc.NewGrpcHandler(service.User),
	}
}
