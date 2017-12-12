package main

import (
	"fmt"
	"net/http"
)

func Exp(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("workingffffffffffff"))
}

func main() {
	fmt.Println("testingn https")

	http.HandleFunc("/", Exp)

	err := http.ListenAndServeTLS(":444", "server.crt", "server.key", nil)
	if err != nil {
		fmt.Println("big error", err)
	}
}
