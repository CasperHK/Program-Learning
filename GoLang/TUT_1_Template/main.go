package main

import (
	"log"
	"os"
	"text/template"
)

func main() {

	// Link up the .gohtml file
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil { // Check Error
		log.Fatalln(err)
	}

	//
	nf, err := os.Create("index.html")
	if err != nil {
		log.Println("error creating file", err)
	}
	defer nf.Close() // This line will be executed after the main block.

	//
	err = tpl.Execute(nf, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
