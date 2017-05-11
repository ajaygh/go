package main

//Contains utility functions and types

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
)

//Parses config and saves the in Conf on successful reading
func readConfig(confFile, confPath, confType string) {
	viper.SetConfigName(confFile)
	viper.AddConfigPath(confPath)
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error in reading config file. Exiting ", err.Error())
		os.Exit(1)
	}

	//get all config data into a conf struct
	Conf = map[string]string{
		"scan_success":   viper.GetString("common.scan_success"),
		"scan_failure":   viper.GetString("common.scan_failure"),
		"sort_success":   viper.GetString("common.sort_success"),
		"sort_failure":   viper.GetString("common.sort_failure"),
		"inst_name":      viper.GetString(confType + ".inst_name"),
		"cont":           viper.GetString(confType + ".cont"),
		"scan_url":       viper.GetString(confType + ".scan_url"),
		"sort_url":       viper.GetString(confType + ".sort_url"),
		"feedback_url":   viper.GetString(confType + ".feedback_url"),
		"wgh_url":        viper.GetString(confType + ".wgh_url"),
		"event_url":      viper.GetString(confType + ".event_url"),
		"fault_code_url": viper.GetString(confType + ".fault_code_url"),
		"image_url":      viper.GetString(confType + ".image_url"),
		"initConfig_url": viper.GetString(confType + ".initConfig_url"),
	}
	for k, v := range Conf {
		fmt.Printf("%s : %s\n", k, v)
	}
}

//Sets header as required
func SetHeader(req *http.Request) {
	req.Header.Set("Accept", "Application/json")
	req.Header.Set("Content-Type", "Application/json")
	req.Header.Set("Charset", "utf-8")
}

func checkError(eType, errStr string, err error) {
	if err != nil {
		switch eType {
		case "fatal":
			log.Fatalln("fatal:", errStr, err)
			return
		case "panic":
			log.Fatalf("%s: %s unrecoverable. Exiting!!!\n", errStr, err.Error())
			os.Exit(1)
		}
	}
}

func makeRequest(data interface{}, method, url string) *http.Request {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(data)
	req, err := http.NewRequest(method, url, b)
	checkError("fatal", "request creation failed", err)
	SetHeader(req)
	return req
}
