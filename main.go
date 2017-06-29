package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	const PORT = ":8080"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var filePath string
		var contentType string

		if r.URL.Path == "/" {
			filePath = "public/index.html"
		} else {
			filePath = "public" + r.URL.Path
		}

		file, err := os.Open(filePath)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
		}

		defer file.Close()

		switch {
		case strings.HasSuffix(r.URL.Path, "css"):
			contentType = "text/css"
		case strings.HasSuffix(r.URL.Path, "html"):
			contentType = "text/html"
		case strings.HasSuffix(r.URL.Path, "ico"):
			contentType = "image/ico"
		}
		w.Header().Add("Content-Type", contentType)
		io.Copy(w, file)
	})

	http.ListenAndServe(PORT, nil)

}
