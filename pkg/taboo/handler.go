package taboo

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
				if e, ok := r.(*Exception); ok {
					h.catch(e)
				} else {
					h.catch(fromPanic(r.(error)))
				}
			}
		}()
	}

	h.try()
}
