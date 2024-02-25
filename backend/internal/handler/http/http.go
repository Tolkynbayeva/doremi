package http

type Http struct {
	http IHttpInterface
}

type IHttpInterface interface {
	// methods
}

func NewHttpHandler(http IHttpInterface) *Http {
	return &Http{
		http: http,
	}

}
