package serve

import (
	"fmt"
	"net/http"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/vicmanbrile/moneyGolang/serve/handlers"
	"github.com/vicmanbrile/moneyGolang/serve/middlewares/auth"
)

func GoServer() {
	PORT := ":8080"

	ApiSubdomain := "api.localhost"
	AssetsSubdomain := "assets.localhost"
	Auth0Subdomain := "auth.localhost"

	{ // API Rest
		http.HandleFunc(ApiSubdomain+"/", handlers.ShowCredits)
	}

	{ // Auth0
		http.HandleFunc(Auth0Subdomain+"/public", auth.AllAcces)

		http.Handle(Auth0Subdomain+"/privade", auth.EnsureValidTolken()(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// CORS Headers.
				w.Header().Set("Access-Control-Allow-Credentials", "true")
				w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
				w.Header().Set("Access-Control-Allow-Headers", "Authorization")

				w.Header().Set("Content-Type", "application/json")

				tocken := r.Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)

				claims := tocken.CustomClaims.(*auth.CustomClaims)

				if !claims.HasScope("read:messages") {
					w.WriteHeader(http.StatusForbidden)
					w.Write([]byte(`{"message":"Alcanze no logrado."}`))
					return
				}
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{"message":"Hola a un punto publico! Necesitas acceder para ver esto."}`))
			}),
		))
	}

	{ // Assets Request
		http.Handle(AssetsSubdomain+"/", http.StripPrefix("/", http.FileServer(http.Dir("./serve/assets"))))
	}

	{
		http.HandleFunc("/", handlers.DocHandler)
	}

	fmt.Println("Server listing... http:localhost" + PORT)
	http.ListenAndServe(PORT, nil)
}
