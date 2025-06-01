package handlers

import (
	"fmt"
	"net/http"

	"github.com/bllol/go-11_hw/3-task/config"
)

func Debug(w http.ResponseWriter, r *http.Request) {
	cfg := config.AppConfig
	fmt.Fprintf(w, "DEBUG INFO:\nPORT=%s\nAPP_NAME=%s\nVERSION=%s\nUSER_NAME=%s\nGREETING_MSG=%s\nADMIN_NAME=%s\nADMIN_EMAIL=%s\n", cfg.Port, cfg.AppName, cfg.Version, cfg.UserName, cfg.GreetingMsg, cfg.AdminName, cfg.AdminEmail)
}
