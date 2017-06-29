package main

import (
	"net/http"
)

func main() {
	const PORT string = ":8080"
	http.ListenAndServe(PORT, http.FileServer(http.Dir("public")))
}
