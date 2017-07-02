package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	const PORT string = ":8080"

	templates := hydrateTemplates()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requestedFile := r.URL.Path[1:]
		t := templates.Lookup(requestedFile + ".html")
		if t != nil {
			err := t.Execute(w, nil)
			if err != nil {
				log.Println(err)
			}

		} else {
			println("something went wrong")
			w.WriteHeader(http.StatusNotFound)
		}
	})

	http.Handle("/img/", http.FileServer(http.Dir("public")))
	http.Handle("/css/", http.FileServer(http.Dir("public")))

	http.ListenAndServe(PORT, nil)

}

func hydrateTemplates() *template.Template {
	result := template.New("templates")
	const basePath string = "templates"

	template.Must(result.ParseGlob(basePath + "/*.html"))
	return result
}
