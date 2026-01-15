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

func (mngr *Manager) Use(middlewares ...Middleware) {
	mngr.globalMiddlewares = append(mngr.globalMiddlewares, middlewares...)
}

func (mngr *Manager) With(next http.Handler, middlewares ...Middleware) http.Handler {
	n := next

	for _, middleware := range middlewares {
		n = middleware(n)
	}

	for _, globalMiddleware := range mngr.globalMiddlewares {
		n = globalMiddleware(n)
	}

	return n
}
