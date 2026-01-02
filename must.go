package must

import (
	"fmt"
	"os"
)

type Controller[T any] interface {
	Recover(err error) T
}

func Do[T any](c Controller[T]) func(func() (T, error)) T {
	return func(f func() (T, error)) T {
		v, err := f()
		if err != nil {
			v = c.Recover(err)
		}
		return v
	}
}

func Recover[T any](c Controller[T]) func(func() error) {
	return func(f func() error) {
		if err := f(); err != nil {
			c.Recover(err)
		}
	}
}

func Have[T any](c Controller[T]) func(T, error) T {
	return func(t T, err error) T {
		return Do(c)(func() (T, error) {
			return t, err
		})
	}
}

func ControllerExits[T any](code int) Controller[T] {
	return controllerExits[T]{code: code}
}

type controllerExits[T any] struct {
	code     int
	fallback T
}

func (c controllerExits[T]) Recover(err error) T {
	fmt.Println(err)
	os.Exit(c.code)
	return c.fallback
}
