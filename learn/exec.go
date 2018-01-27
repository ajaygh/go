package main

import (
	//"bytes"
	"fmt"
	"log"
	"os/exec"
)

func main() {
	res, err := exec.Command("go", "build", "-o", "webb", "-i", "-v", "github.com/ajaygh/web").Output()
	//cmd.Stdin = strings.NewReader("some input")
	//var out bytes.Buffer
	//cmd.Stdout = &out
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("res ", string(res))
	}
}
