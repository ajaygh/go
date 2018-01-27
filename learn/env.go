package main

import "fmt"
import "os"

//import "strings"

func main() {
	//	os.Setenv("PORT", "4444")
	fmt.Println(os.Getenv("PORT"))
	if port, ok := os.LookupEnv("PORT"); !ok {
		fmt.Println("port not found")
	} else {
		fmt.Println("port = ", port)
	}

	fmt.Println()
	//	for _, e := range os.Environ() {
	//		pair := strings.Split(e, "=")
	//		fmt.Println(pair[0], pair[1])
	//	}
}
