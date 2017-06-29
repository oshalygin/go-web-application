package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	const PORT = ":8080"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		filePath := "public" + r.URL.Path
		file, err := os.Open(filePath)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
		}

		defer file.Close()
		io.Copy(w, file)
	})

	http.ListenAndServe(PORT, nil)

}
