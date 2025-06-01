package main

import (
	"net/http"

	"github.com/bllol/go-11_hw/2-task/config"
	"github.com/bllol/go-11_hw/2-task/routes"
)

func main() {
	config.LoadConfig()
	cfg := config.AppConfig

	routes.RegisterRoutes()

	http.ListenAndServe(":"+cfg.Port, nil)
}
