package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler // type or signature

/*
struct : struct look like a template for object or model build
*/
type Manager struct {
	globalMiddlewares []Middleware
}

/*
pointer: pointer declare return before value type (*) and return value (&) give
empty array: make(value, value length)
*/
func NewManager() *Manager {
	mngr := Manager{
		globalMiddlewares: make([]Middleware, 0),
	}

	return &mngr
}

func (mngr *Manager) With(middlewares ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		n := next

		for i := len(middlewares) - 1; i >= 0; i-- {
			middleware := middlewares[i]
			n = middleware(n)
		}

		return n
	}
}
