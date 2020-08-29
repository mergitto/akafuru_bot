package gcp_function_subscriber

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/nlopes/slack"
)

var verificationToken string

const (
	apiKey = "pae7leipaithopo5achaePh1eiwee3feju2aili8Eijua8ca"
)

func init() {
	verificationToken = os.Getenv("VERIFICATION_TOKEN")
}

func HelloCommand(w http.ResponseWriter, r *http.Request) {
	if os.Getenv("API_KEY") != apiKey {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// スラッシュコマンドのリクエストをパースする。
	s, err := slack.SlashCommandParse(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Slack から来るリクエストに付与される ValidateToken をチェックする。
	if !s.ValidateToken(verificationToken) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	switch s.Command {
	case "/hello":
		params := &slack.Msg{ResponseType: "in_channel", Text: "こんにちは、<@" + s.UserID + ">さん"}
		b, err := json.Marshal(params)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	default:
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
