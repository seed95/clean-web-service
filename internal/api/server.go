package api

type (
	HttpServer interface {
		Start(port int) error
	}
)
