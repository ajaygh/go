package main

import (
	"fmt"

	textmagic "github.com/textmagic/textmagic-rest-go"
)

func main() {
	client := textmagic.NewClient("<USERNAME>", "<APIV2_KEY>")

	params := map[string]string{
		"phones": "9110962454",
		"text":   "Hello from TextMagic Go",
	}
	message, err := client.CreateMessage(params)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(message.ID)
	}
}
