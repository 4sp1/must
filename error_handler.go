package must

type ErrorHandlerFn func(func() error)
type ErrorValueHandlerFn func(error)

type ErrorHandler interface {
	Handle(err error)
}

type errorHandler[T any] struct {
	c Controller[T]
}

func (h errorHandler[T]) Handle(err error) {
	h.c.Fallback(err)
}

func HandlerOf[T any](c Controller[T]) ErrorHandler {
	return errorHandler[T]{c}
}
