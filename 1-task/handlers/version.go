package handlers

import (
	"fmt"
	"github.com/bllol/go-11_hw/1-task/config"
	"net/http"
)

func Version(w http.ResponseWriter, r *http.Request) {
	cfg := config.AppConfig
	fmt.Fprintf(w, "App: %s\nVersion: %s", cfg.AppName, cfg.Version)
}
