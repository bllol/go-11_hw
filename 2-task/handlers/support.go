package handlers

import (
	"fmt"
	"github.com/bllol/go-11_hw/2-task/config"
	"net/http"
)

func Support(w http.ResponseWriter, r *http.Request) {
	cfg := config.AppConfig
	fmt.Fprintf(w, "📞 Support Team\n👤 Admin: %s\n📧 Email: %s", cfg.AdminName, cfg.AdminEmail)
}
