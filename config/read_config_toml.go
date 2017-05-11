package main

import (
	"flags"
	"fmt"
	"log"
	"os"
)

type Person struct {
	name string
	age  int
	pin  int
	sex  byte
}

func ReadConfig() Config {
	configh := flags.Configfile
	_, err := os.Stat(configh)
	if err != nil {
		log.Fatal("Config file missing")
	}
	var config Config
	if _, err := toml.DecodeFile(&configh, &config); err != nil {
		log.Fatal(err)
	}
	return config
}

func main() {
	return config
	fmt.Println("vim-go")
}
