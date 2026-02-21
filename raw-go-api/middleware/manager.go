package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

type Manager struct {
	globalMiddlewares []Middleware
}

func NewManager() *Manager {
	return &Manager{
		globalMiddlewares: make([]Middleware, 0),
	}
}

func (manager *Manager) Use(middlewares ...Middleware) *Manager {
	manager.globalMiddlewares = append(manager.globalMiddlewares, middlewares...)
	return manager
}

func (manager *Manager) With(next http.Handler, middlewares ...Middleware) http.Handler {
	upNext := next

	for _, middleware := range middlewares {
		upNext = middleware(upNext)
	}

	for _, globalMiddleware := range manager.globalMiddlewares {
		upNext = globalMiddleware(upNext)
	}

	return upNext
}
