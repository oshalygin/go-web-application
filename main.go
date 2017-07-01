package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	templateString := `Foobar Stand Supply`
	t, err := template.New("title").Parse(templateString)

	if err != nil {
		log.Println(err)
	}

	err = t.Execute(os.Stdout, nil)
	if err != nil {
		log.Println(err)
	}

}
