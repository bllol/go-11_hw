package routes

import (
	"github.com/bllol/go-11_hw/3-task/handlers"
	"net/http"
)

func RegisterRoutes() {
	http.HandleFunc("/debug", handlers.Debug)
}
