package taboo

import "runtime"

type handler struct {
	try     func()
	catch   func(e *Exception)
	finally func()
}

func Try(f func()) *handler {
	return &handler{
		try: f,
	}
}

func (h *handler) Catch(f func(e *Exception)) *handler {
	h.catch = f
	return h
}

func (h *handler) Finally(f func()) *handler {
	h.finally = f
	return h
}

func (h *handler) Do() {
	if h.finally != nil {
		defer h.finally()
	}

	if h.catch != nil {
		defer func() {
			if r := recover(); r != nil {
				switch e := r.(type) {
				case *Exception:
					h.catch(e)
				case runtime.Error:
					h.catch(fromError(e))
				case interface{}:
					h.catch(fromInterface(e))
				}
			}
		}()
	}

	h.try()
}
