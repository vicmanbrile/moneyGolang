package middlewares

import (
	"net/http"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

func Logging() Middleware {
	return func(NextHandler http.HandlerFunc) http.HandlerFunc {
		handler := func(w http.ResponseWriter, r *http.Request) {
			Cookie, _ := r.Cookie("Profile")

			if Cookie != nil {
				NextHandler(w, r)
			}

		}

		return handler

	}
}

/*

Example to create a Middleware

func NewMiddleware() Middleware {

	return func(NextHandler http.HandlerFunc) http.HandlerFunc {

		handler := func(w http.ResponseWriter, r *http.Request) {

			... White partther for middlerwares

			NextHandler(w, r)

		}

		return handler

	}

}


func Auth() Middleware {

	return func(NextHandler http.HandlerFunc) http.HandlerFunc {

		handler := func(w http.ResponseWriter, r *http.Request) {

			// ... White partther for middlerwares

			c, _ := r.Cookie("Profile")

			if c.Value != "" {
				Form := template.New("Form")

				Form.Parse(templates.FormCredit)

				Success := true
				Form.Execute(w, Success)
			}

		}

		return handler
	}

}

*/
