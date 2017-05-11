package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

const (
	scan      = "/api/v1/controller/c1/shipment/scan/"
	sort      = "/api/v1/controller/c1/shipment/sort/"
	event     = "/api/v1/controller/c1/event/event/"
	operation = "/api/v1/controller/c1/operation/"
)

type ScanSnd struct {
	ChuteId string `json:"chuteid"`
}

type Event struct {
	Code string `json:"event_code"`
}

type Operation struct {
	Code string `json:"operation_type"`
}

//Scan receiver
type Scan struct {
	IcrId        string `json:"icrid"`
	JobId        string `json:"jobid"`
	CasketId     string `json:"casketid"`
	Width        string `json:"breadth"`
	Length       string `json:"length"`
	Height       string `json:"height"`
	RealVol      string `json:"real_volume"`
	BoxVol       string `json:"box_volume"`
	VmsStatus    string `json:"vms_status_code"`
	WeightStatus string `json:"weight_status_code"`
	Weight       string `json:"weight"`
	InputNo      string `json:"inductid"`
	ImageId      string `json:"imageid"`
	Uuid         string `json:"uid"`
	Barcode      string `json:"barcode"`
	ScanStatus   string `json:"scan_status_code"`
}

//Receive sort data
type Sort struct {
	JobId      string `json:"jobid"`
	Uuid       string `json:"uid"`
	SortStatus string `json:"sort_status_code"`
	ChuteId    string `json:"chuteid"`
}

func checkBody(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
}

func checkDecodeError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}

func handleScan(w http.ResponseWriter, r *http.Request) {
	checkBody(w, r)
	var sd Scan
	err := json.NewDecoder(r.Body).Decode(&sd)
	checkDecodeError(w, err)

	fmt.Println("Decoded successfully", sd)
	chute := rand.Intn(250)
	ss := ScanSnd{ChuteId: strconv.Itoa(chute)}
	err = json.NewEncoder(w).Encode(&ss)
	fmt.Println(time.Now(), ss, "sent api scan")
}

func handleSort(w http.ResponseWriter, r *http.Request) {
	checkBody(w, r)
	var sd Sort
	err := json.NewDecoder(r.Body).Decode(&sd)
	checkDecodeError(w, err)

	fmt.Println("Decoded successfully", sd)

	ss, _ := json.Marshal("detail : ok")
	err = json.NewEncoder(w).Encode(&ss)
	fmt.Println(time.Now(), ss, "sent api scan")
}

func handleEvent(w http.ResponseWriter, r *http.Request) {
	checkBody(w, r)
	var event Event
	err := json.NewDecoder(r.Body).Decode(&event)
	checkDecodeError(w, err)

	now := time.Now()
	n := rand.Intn(30)
	time.Sleep(time.Second * time.Duration(n))
	now1 := time.Now()
	fmt.Println(now1, "Successful decode Waited seconds", now.After(now1))
	fmt.Fprintf(w, "res sent after %d", n)
}

func handleOperation(w http.ResponseWriter, r *http.Request) {
	checkBody(w, r)
	var opn Operation
	err := json.NewDecoder(r.Body).Decode(&opn)
	checkDecodeError(w, err)
	//Process response
	now := time.Now()
	fmt.Println(now, " response sent for ", opn.Code)
	fmt.Fprintf(w, "time %s Res %s received", now.String(), opn.Code)
}

func main() {
	http.HandleFunc(scan, handleScan)
	http.HandleFunc(event, handleEvent)
	http.HandleFunc(sort, handleSort)
	http.HandleFunc(operation, handleOperation)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatalln("ListenAndServe: ", err)
	}
	fmt.Println("vim-go")
}
