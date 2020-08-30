package handler

import (
	"encoding/json"
	"net/http"

	"github.com/nlopes/slack"
)

func writeBody(w http.ResponseWriter, buf []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(buf)
}

func slackMsg2Marshal(m *slack.Msg) ([]byte, error) {
	b, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return b, nil
}
