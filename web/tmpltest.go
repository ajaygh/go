package main

import (
	"html/template"
	"os"
)

type Friend struct {
	Fname string
}

type Person struct {
	FirstName string
	LastName  string
	Emails    []string
	Friends   []*Friend
}

func main() {
	f1 := Friend{"Tom"}
	// f2 := Friend{"John"}
	// f3 := Friend{"Michael"}
	// emails := []string{"d12@gmail.com", "d32@gmail.com"}
	// person := Person{"Sachin", "Tendulkar", emails, []*Friend{&f1, &f2, &f3}}
	t := template.New("test")
	t, _ = t.Parse("My friend is {{.Fname}}\n")
	// t, _ = t.Parse(`Your name {{.FirstName}} {{.LastName}}
	// 				Your emails are
	// 				{{range .Emails}}
	// 					an email {{.}}
	// 				{{end}}
	// 				Your friends are
	// 				{{with .Friends}}
	// 				{{range .}}
	// 				 	my friend is {{.Fname}}
	// 				{{end}}
	// 				{{end}}
	// 		`)
	//	Name := "ajay"
	//t.Execute(os.Stdout, Person{FirstName: Name, LastName: "Kumar"})
	t.Execute(os.Stdout, f1)
}
