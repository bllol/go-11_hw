package routes

import (
	"github.com/bllol/go-11_hw/1-task/handlers"
	"net/http"
)

func RegisterRoutes() {
	http.HandleFunc("/version", handlers.Version)
}
