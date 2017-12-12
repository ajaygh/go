package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func SetHeader(req *http.Request) {
	req.Header.Set("Content-Type", "application/vnd.onem2m-res+json;ty=4")
	req.Header.Set("X-M2M-Origin", "C5F414079-304954fa")
	req.Header.Set("X-M2M-RI", "12444328")
	//	req.Header.Set("Authorization", "Basic QzVGNDE0MDc5LTMwNDk1NGZhOnRlc3RAMTIz")
	req.Header.Set("Accept", "application/vnd.onem2m-res+json;")
}

func main() {
	//getLatestPayload()
	//getAllAssets()
	//getIndex()
	sendDownlinkPayload()
}

type Payload struct {
	Deveui string `json:"deveui"`
}

func sendDownlinkPayload() {
	indata := `{"m2m:cin": {"ty":4,	"cnf":"text/plain:0","cs":300,"con":"{\"payload_dl\":{\"deveui\":\"0004a30b001ba065\",\"port\":2,\"confirmed\":true,\"data\":\"ODAwMA==\",\"on_busy\":\"fail\",\"tag\":\"98861544465w\"}}"}}`
	b := []byte(indata)
	fmt.Println(string(b))
	//json.NewEncoder(c).Encode(pl)
	req, err := http.NewRequest("POST", "http://168.87.87.213:8080/davc/m2m/HPE_IoT/0004a30b001ba065/DownlinkPayload", bytes.NewBuffer(b))
	//req, err := http.NewRequest("POST", "http://168.87.87.213:8080/davc/m2m/HPE_IoT/0004a30b001b9231/DownlinkPayload", nil)

	chkErr(err, "err in req: ")
	SetHeader(req)
	//req.SetBasicAuth("gauravkumar.sinha@havells.com", "Gaurav@3093")
	req.SetBasicAuth("C5F414079-304954fa", "test@123")

	client := &http.Client{Timeout: time.Second * 15}
	resp, err := client.Do(req)
	chkErr(err, "err in resp: ")
	var test interface{}
	json.NewDecoder(resp.Body).Decode(&test)
	fmt.Println("received data ", test)
	resp.Body.Close()
}

func getLatestPayload() {
	req, err := http.NewRequest("GET", "http://168.87.87.213:8080/davc/m2m/HPE_IoT/0004a30b001b02da/default/latest", nil)
	chkErr(err, "err in req: ")
	SetHeader(req)
	client := &http.Client{Timeout: time.Second * 25}
	resp, err := client.Do(req)
	chkErr(err, "err in resp: ")
	var test interface{}
	json.NewDecoder(resp.Body).Decode(&test)
	fmt.Println("received data ", test)
	resp.Body.Close()
}

func getAllAssets() {
	req, err := http.NewRequest("GET", "http://168.87.87.210:8080/dsm/wsi/asset/getAllAssets/gauravkumar.sinha@havells.com/", nil)
	chkErr(err, "err in req: ")
	req.Header.Set("Authorization", "Basic Z2F1cmF2a3VtYXIuc2luaGFAaGF2ZWxscy5jb206R2F1cmF2QDMwOTM=")
	client := &http.Client{Timeout: time.Second * 15}
	resp, err := client.Do(req)
	chkErr(err, "err in resp: ")
	buf, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("received data ", string(buf))
	resp.Body.Close()
}

func getIndex() {
	resp, err := http.Get("http://34.210.167.50:9000/")
	//resp, err := http.Get("https://curl.haxx.se")
	chkErr(err, "err in resp: ")
	buf, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("received data ", string(buf))
	resp.Body.Close()
}

func chkErr(err error, msg string) {
	if err != nil {
		fmt.Println(msg, err)
		os.Exit(1)
	}
}
