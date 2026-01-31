package app

import (
	"fmt"
	"net/http"
	"time"
	"urlshortner/controller"
	config "urlshortner/db"
	"urlshortner/repository"
	"urlshortner/routers"
	"urlshortner/services"
)

type Config struct {
	Address string
}

type Application struct {
	Config
}

func NewConfig(address string) Config {
	return Config{
		Address: address,
	}
}
func NewApplication(config Config) *Application {
	return &Application{
		Config: config,
	}
}

func (app *Application) Run() error {
	db, err := config.SetUpDb()
	if err != nil {
		return err 
	}
	fmt.Println("Starting server on",app.Config.Address)
	ur:=repository.NewUrlRepository(db)
	us:=services.NewUrlServiceImpl(ur)
	uc:=controller.NewUrlController(us)
	uR:=routers.NewUrlRouter(uc)
	server:=&http.Server{
		Addr: app.Config.Address,
		Handler: routers.SetUpRouter(uR),
		ReadTimeout: 10 *time.Second,
		WriteTimeout: 10 *time.Second,
	}
	return server.ListenAndServe()
}