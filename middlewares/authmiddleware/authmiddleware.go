package authmiddleware

import (
	"errors"
	"net/http"

	"github.com/juliotorresmoreno/proxy-logger/helpers"
)

// AuthMiddleware .
func AuthMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		if authorization[:6] != "Bearer" {
			helpers.HandleHTTPError(w, r, errors.New("Unauthorized"), 401)
			return
		}
		jwtToken := authorization[7:]
		_, err := helpers.ValidateJwtToken(jwtToken)
		if helpers.HandleHTTPError(w, r, err, 401) {
			return
		}
		handler.ServeHTTP(w, r)
	})
}
