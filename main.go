package main

import (
	"fmt"
	"github.com/ashwanthkumar/slack-go-webhook"
	"github.com/yoshimitsuEgashira/random-roles/payload"
	"os"
)

func main() {
	fmt.Println("Start Clearing Duty ʕ ◔ ϖ ◔ʔ ==Go")
	p := payload.MakeAttachment()
	fmt.Println("Requesting Slack API ...")
	postSlack(p)
}

func postSlack(p slack.Payload) {
	u := os.Getenv("URL")
	err := slack.Send(u, "", p)
	if len(err) > 0 {
		fmt.Printf("Error: %s\n", err)
	}
	fmt.Println("Notify to your slack channel completed !!")
}
