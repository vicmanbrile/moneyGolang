package serve

import (
	"html/template"
	"net/http"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func Home(w http.ResponseWriter, r *http.Request) {

	/*var Layout string = `
	<h1>{{.PageTitle}}</h1>
	<ul>
		{{range .Todos}}
			{{if .Done}}
				<li class="done">{{.Title}}</li>
			{{else}}
				<li>{{.Title}}</li>
			{{end}}
		{{end}}
	</ul>
	`*/

	tmpl := template.New("Home")

	data := TodoPageData{
		PageTitle: "My TODO list",
		Todos: []Todo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: true},
		},
	}

	tmpl.Execute(w, data)
}
