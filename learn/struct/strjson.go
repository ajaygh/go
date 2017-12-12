package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type DL struct {
	Data string `json:"data"`
	Fcnt int    `json:"fcnt"`
}

type UL struct {
	DataFrame string `json:"dataFrame"`
	Fcnt      int    `json:"fcnt"`
}

func main() {
	fmt.Println("json decode testing")
	sdl := `{"payload_dl":{"data":"hhhhhgggg","fcnt":5}}`
	sul := `{"payload_ul":{"dataFrame":"hhhhhulll","fcnt":5}}`
	var dl DL
	var bsdl []byte
	bsdl = []byte(sul)
	if strings.Contains(string(bsdl), "payload_dl") {
		tmp := struct {
			PayloadDl DL `json:"payload_dl"`
		}{}
		json.NewDecoder(strings.NewReader(sdl)).Decode(&tmp)
		fmt.Println("tmp is ", tmp)
		err := json.Unmarshal(bsdl, &tmp)
		if err != nil {
			fmt.Println("error is", err)
			return
		}
		fmt.Println("dl is ", tmp)
		json.NewEncoder(os.Stdout).Encode(dl)
	} else if strings.Contains(string(bsdl), "payload_ul") {
		fmt.Println("uplink found")
		tmp := struct {
			PayloadDl UL `json:"payload_ul"`
		}{}
		json.NewDecoder(strings.NewReader(sul)).Decode(&tmp)
		fmt.Println("uplink is", tmp.PayloadDl)
	}
}
