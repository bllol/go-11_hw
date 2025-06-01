package main

import (
	"github.com/bllol/go-11_hw/3-task/config"
	"github.com/bllol/go-11_hw/3-task/routes"
	"net/http"
)

func main() {
	config.LoadConfig()
	cfg := config.AppConfig

	routes.RegisterRoutes()

	http.ListenAndServe(":"+cfg.Port, nil)
}
