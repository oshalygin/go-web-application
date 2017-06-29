package main

import (
	"net/http"
)

func main() {
	const PORT = ":8080"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "public"+r.URL.Path)
	})

	http.ListenAndServe(PORT, nil)

}
