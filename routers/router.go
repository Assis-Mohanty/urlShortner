package routers

import "github.com/go-chi/chi/v5"

type Router interface {
	Register(r *chi.Mux)
}

func SetUpRouter(routes ...Router) *chi.Mux{
	router:=chi.NewRouter()
	for _, route := range routes {
		route.Register(router)
	}
	return router
}
