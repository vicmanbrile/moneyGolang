package middlewares

import (
	"crypto/sha256"
	"crypto/subtle"
	"net/http"
)

func Loggin(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		user, pass, ok := r.BasicAuth()

		if ok {

			usernameHash := sha256.Sum256([]byte(user))
			passwordHash := sha256.Sum256([]byte(pass))
			expectedUsernameHash := sha256.Sum256([]byte("vicmanbrile"))
			expectedPasswordHash := sha256.Sum256([]byte("Fenian_135"))

			usernameMatch := (subtle.ConstantTimeCompare(usernameHash[:], expectedUsernameHash[:]) == 1)
			passwordMatch := (subtle.ConstantTimeCompare(passwordHash[:], expectedPasswordHash[:]) == 1)
			if usernameMatch && passwordMatch {

				next.ServeHTTP(w, r)
				return
			}

		}

		w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)

		http.Error(w, "Unauthorized", http.StatusUnauthorized)

	})
}
