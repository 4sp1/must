package must

func Do[T any](c Controller[T]) ControllerDoFn[T] {
	return func(f func() (T, error)) T {
		v, err := f()
		if err != nil {
			v = c.Fallback(err)
		}
		return v
	}
}

func Have[T any](c Controller[T]) ControllerHaveFn[T] {
	return func(t T, err error) T {
		return Do(c)(func() (T, error) {
			return t, err
		})
	}
}

func Handle(h ErrorHandler) ErrorHandlerFn {
	return func(f func() error) {
		if err := f(); err != nil {
			h.Handle(err)
		}
	}
}

func HandleError(h ErrorHandler) ErrorValueHandlerFn {
	return func(err error) {
		if err != nil {
			h.Handle(err)
		}
	}
}
