package routes

import (
	"github.com/bllol/go-11_hw/2-task/handlers"
	"net/http"
)

func RegisterRoutes() {
	http.HandleFunc("/support", handlers.Support)
}
