package handler

import (
	"net/http"

	"github.com/nlopes/slack"
)

func Hello(w http.ResponseWriter, s slack.SlashCommand) {
	params := &slack.Msg{ResponseType: "in_channel", Text: "こんにちは、<@" + s.UserID + ">さん"}
	b, err := slackMsg2Marshal(params)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	writeBody(w, b)
}
