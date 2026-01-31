package main

import (
	"fmt"
	"urlshortner/app"
	config "urlshortner/db"
)

func main() {
	fmt.Println("kasndkads")
	config.Load()
	cfg:=app.NewConfig(":3000")
	app:=app.NewApplication(cfg)
	app.Run()
}