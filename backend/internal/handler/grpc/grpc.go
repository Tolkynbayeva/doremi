package grpc

import "doremi/internal/service"

type Grpc struct {
	service service.Service
}

func NewGrpcHandler(service service.Service) *Grpc {
	return &Grpc{
		service: service,
	}
}
