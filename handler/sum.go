package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/nlopes/slack"
)

func Sum(w http.ResponseWriter, s slack.SlashCommand) {
	text := strings.ReplaceAll(s.Text, "　", " ")
	strs := strings.Split(text, " ")
	sum := 0
	for _, str := range strs {
		num, err := strconv.Atoi(str)
		if err != nil {
			log.Printf("failed parse int: [%s]", str)
			writeBody(w, []byte("数値を入力してください"))
			return
		}
		sum += num
	}

	params := &slack.Msg{ResponseType: "in_channel", Text: fmt.Sprintf("合計は %d ですニャー :cat2:", sum)}
	b, err := json.Marshal(params)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	writeBody(w, b)
}
