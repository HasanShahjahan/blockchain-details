package middleware

import (
	auth "github.com/NuriCareers/Shahjahan-Miah-coding-challenge/internal/auth"
	"github.com/NuriCareers/Shahjahan-Miah-coding-challenge/internal/config"
	"github.com/NuriCareers/Shahjahan-Miah-coding-challenge/internal/logger"
	u "github.com/NuriCareers/Shahjahan-Miah-coding-challenge/internal/utils"
	"strconv"
)

import (
	"net/http"
)

const (
	logTag = "Middleware"
)

//SetMiddlewareJSON ...
func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

//SetMiddlewareAuthentication ...
func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logging.Info(logTag, "JWT authentication token :"+strconv.FormatBool(config.Config.IsJwtEnabled))
		if config.Config.IsJwtEnabled {
			err := auth.TokenValid(r)
			if err != nil {
				u.RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
				logging.Error(logTag, "Unauthorized", err)
				return
			}
		}
		next(w, r)
	}
}
