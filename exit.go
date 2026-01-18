package must

import (
	"fmt"
	"os"
)

func ExitHandler(code int) ErrorHandler {
	return HandlerOf(exitController[struct{}]{code: code})
}

func ExitController[T any](code int) Controller[T] {
	return exitController[T]{code: code, exitFunc: os.Exit}
}

type exitController[T any] struct {
	code     int
	fallback T              // only to satisfy exitController.Fallback return
	exitFunc func(code int) // makes exit behavior testable
}

func (c exitController[T]) Fallback(err error) T {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(c.code)
	return c.fallback
}
