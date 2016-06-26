package main

import (
	"log"
	"os"
	"text/template"
)

type person struct {
	Name string
	Age  int
}

func main() {
	p1 := person{
		Name: "James BOnd",
		Age:  42,
	}

	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, p1)
	if err != nil {
		log.Fatalln(err)
	}
}
