package main

/*
This package generates random number of words. This number of words is taken as command line argument.
If no argument is provided then it generates only one word.E.g
./randword
apple
./randword 4
apple
mango
where
now
*/

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// at max only one argument required
	if len(os.Args) > 2 {
		log.Println("Extra arguments received.Please provide at max 1 argument")
		os.Exit(1)
	}

	genLen := 1
	if len(os.Args) == 2 {
		tmpLen, err := strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatalln("err in argument value:", err)
		}
		genLen = tmpLen
		//log.Println("generator length is ", genLen)
	}

	// read words and store them in slice
	rwin, err := ioutil.ReadFile("random_word.txt")
	if err != nil {
		log.Fatalln("err in reading random_word.txt file:", err)
	}
	words := strings.Split(string(rwin), "\n")
	// for i := range words {
	// 	log.Printf("index %d word %s\n", i, words[i])
	// }

	maxLen := len(words)
	if genLen > maxLen {
		log.Fatalf("argument length %d is greater than number %d of available words\n", genLen, maxLen)
	}
	// Seed the source
	rand.Seed(time.Now().UnixNano())
	randIdx := rand.Intn(maxLen - genLen + 1)
	for i := 0; i < genLen; i++ {
		fmt.Println(words[randIdx+i])
	}
}
