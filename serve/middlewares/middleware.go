package middlewares

import "net/http"

type Middleware func(http.HandlerFunc) http.HandlerFunc

func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {

	for _, m := range middlewares {
		f = m(f)
	}

	return f

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

*/