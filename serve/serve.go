package serve

import (
	"fmt"
	"net/http"

	"github.com/vicmanbrile/moneyGolang/serve/handlers/middlewares"
	"github.com/vicmanbrile/moneyGolang/serve/handlers/router"
)

type Handle func(http.ResponseWriter, *http.Request)
type Server struct {
	Port   string
	Router *router.Router
}

func NuevoServidor(port string) *Server {
	return &Server{
		Port:   port,
		Router: router.NuevoRouter(),
	}
}

func (s *Server) GoServer() {
	AssetsSubdomain := "assets.localhost"

	http.Handle("/", s.Router)
	// Auth0Subdomain := "auth.localhost"

	/*{ // Auth0
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
	}*/

	{ // Assets Request
		http.Handle(AssetsSubdomain+"/", http.StripPrefix("/", http.FileServer(http.Dir("./assets"))))
	}

	fmt.Println("Server listing... http:localhost" + s.Port)
	http.ListenAndServe(s.Port, nil)
}

func (s *Server) Handle(path string, handle http.HandlerFunc, method string, midd ...middlewares.Middleware) http.HandlerFunc {

	RV := &router.RuteValidation{
		HDF:          handle,
		Middlerwares: midd,
	}

	{
		if s.Router.Rules[path] == nil {
			s.Router.Rules[path] = make(map[string]router.RuteValidation)
		}

		s.Router.Rules[path][method] = *RV
	}

	return handle
}
