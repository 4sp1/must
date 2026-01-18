package must

type Controller[T any] interface {
	Fallback(err error) T
}

type ControllerDoFn[T any] func(func() (T, error)) T

type ControllerHaveFn[T any] func(T, error) T
