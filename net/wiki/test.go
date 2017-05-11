package main

import "fmt"

type Page struct {
	Title string
	Body  string
}

func writeTest(p Page) {
	p.Body = fmt.Sprintf("just testing %s", p.Title)
	fmt.Println(p.Body)
}

func main() {
	p := Page{Title: "first"}
	writeTest(p)
	fmt.Printf("Title: %s Body : %s\n", p.Title, p.Body)
}
