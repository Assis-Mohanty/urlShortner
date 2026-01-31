package routers

import (
	"fmt"
	"urlshortner/controller"

	"github.com/go-chi/chi/v5"
)

type UrlRouter struct {
	urlController *controller.UrlController
}

func NewUrlRouter(urlController *controller.UrlController) Router{
	return &UrlRouter{
		urlController: urlController,
	}
}
func(ur *UrlRouter)Register(r *chi.Mux){
	fmt.Println("Registering URL routes")
	r.Get("/url/long-id", ur.urlController.GetIdByLongUrlHandler)
	r.Get("/url/short/{shorturl}", ur.urlController.GetLongUrlHandler)
	r.Post("/url/shorturl",ur.urlController.CreateShortUrlHandler)
	r.Post("/url/longurl",ur.urlController.CreateLongUrl)
}