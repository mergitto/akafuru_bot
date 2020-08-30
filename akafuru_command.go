package gcp_function_subscriber

import (
	"net/http"
	"os"

	"github.com/mergitto/akafuru_bot/handler"
	"github.com/nlopes/slack"
)

var verificationToken string

func init() {
	verificationToken = os.Getenv("VERIFICATION_TOKEN")
}

func AkafuruCommand(w http.ResponseWriter, r *http.Request) {
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
		handler.Hello(w, s)
	case "/sum":
		handler.Sum(w, s)
	default:
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
