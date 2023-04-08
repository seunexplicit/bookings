package middlewares

import (
	"net/http"

	"github.com/seunexplicit/bookings/internal/config"
)

var appConfig *config.AppConfig

func LoadAppConfig(config *config.AppConfig) {
	appConfig = config
}

func LoadSession(next http.Handler) http.Handler {
	return appConfig.Session.LoadAndSave(next)
}
