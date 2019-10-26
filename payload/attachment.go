package payload

import (
	"fmt"
	"github.com/ashwanthkumar/slack-go-webhook"
)

func MakeAttachment() slack.Payload {
	fmt.Println("Making attachment for sending slack ...")
	a := slack.Attachment{}
	a.AddField(slack.Field{Title: "テクラボ福岡 掃除当番", Value: MakeAttachmentText(), Short: false})

	p := slack.Payload{
		// Change the info such as the channel, text and others you want to be notified here.
		Text:        "今週もお掃除の時間がやって参りました！\n本日の当番を発表します ʕ ◔ ϖ ◔ ʔ",
		Username:    "Random Clearing Roles",
		Channel:     "#tech-lab_all",
		IconEmoji:   ":gopher:",
		Attachments: []slack.Attachment{a},
	}
	return p
}
