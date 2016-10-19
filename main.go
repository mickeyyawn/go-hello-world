package main

import "net/http"

//import "github.com/gorilla/mux"
import "html/template"
import "github.com/kataras/iris"

type Controller struct {
}

func main() {

	iris.Get("/", func(ctx *iris.Context) {
		ctx.Write("Hello world %s", ", this is new, from golang...")
	})
	iris.Listen(":8080")

	/*
		port := ":8080"

		r := mux.NewRouter()

		r.HandleFunc("/", root)
		http.Handle("/", r)

		http.ListenAndServe(port, nil)
	*/

}

func root(w http.ResponseWriter, r *http.Request) {

	html := render()

	tmpl, err := template.New("name").Parse(html)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func render() string {

	html := `<!doctype html>
	           <head>
	             <title>hello world!</title>
	           </head>
	           <body>
	             <h1>Hello World from golang!</h1>
	           </body>
           </html>`

	return html

}
