package grpc

type Grpc struct {
	grpc IGrpcInterface
}

type IGrpcInterface interface {
	// methods
}

func NewGrpcHandler(grpc IGrpcInterface) *Grpc {
	return &Grpc{
		grpc: grpc,
	}
}
