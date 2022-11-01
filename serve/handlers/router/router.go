package router

import (
	"net/http"

	"github.com/vicmanbrile/moneyGolang/serve/handlers/middlewares"
)

type RuteValidation struct {
	HDF          http.HandlerFunc
	Middlerwares []middlewares.Middleware
}

type Router struct {
	// ["NombreHandle":["Method":HandlerFunc]]
	Rules map[string]map[string]RuteValidation
}

func NuevoRouter() *Router {
	return &Router{
		Rules: make(map[string]map[string]RuteValidation),
	}
}

func (r *Router) FindHandler(path, method string) (Status int) {

	_, PathExist := r.Rules[path]

	if PathExist {

		_, MethodExist := r.Rules[path][method]

		if MethodExist {
			Status = http.StatusOK
		} else {
			Status = http.StatusMethodNotAllowed
		}

	} else {
		Status = http.StatusNotFound
	}

	return
}

func (r *Router) ServeHTTP(w http.ResponseWriter, rq *http.Request) {
	RV := r.Rules[rq.URL.Path][rq.Method]

	if len(RV.Middlerwares) > 0 {
		for _, m := range RV.Middlerwares {
			RV.HDF = m(RV.HDF)
		}

		RV.HDF(w, rq)
	} else {
		RV.HDF(w, rq)
	}

}
