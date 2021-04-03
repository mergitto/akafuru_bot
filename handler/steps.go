package handler

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/go-resty/resty/v2"
	"github.com/nlopes/slack"
)

type step struct {
	DateMonth string `json:"date_month"`
	Value     int    `json:"value"`
}

type stepsEachMonth struct {
	Steps    []*step `json:"steps"`
	AllSteps int     `json:"all_steps"`
}

func Steps(w http.ResponseWriter, s slack.SlashCommand) {
	msg, err := toMessage(getSteps(w, s.Text), s.Text)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	b, err := slackMsg2Marshal(msg)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	writeBody(w, b)
}

func toMessage(steps *stepsEachMonth, yearMonth string) (*slack.Msg, error) {
	var msg string
	if yearMonth != "" {
		msg = yearMonth
	} else {
		msg = "今月"
	}
	msg += "の歩数をお知らせするニャー！\n"
	for _, step := range steps.Steps {
		t, err := time.Parse("2006-01-02", step.DateMonth)
		if err != nil {
			return nil, err
		}
		month := fmt.Sprintf("%d年%d月%d日", t.Year(), t.Month(), t.Day())
		stepComma := humanize.Comma(int64(step.Value))
		switch {
		case step.Value < 3000:
			msg += fmt.Sprintf("%s: %s歩 (デブへの道へと続いてるニャ…)\n", month, stepComma)
		case step.Value > 10000:
			msg += fmt.Sprintf("%s: %s歩 (そんなに運動してどうするニャ!!)\n", month, stepComma)
		case step.Value > 8000:
			msg += fmt.Sprintf("%s: %s歩 (よく運動したニャ!)\n", month, stepComma)
		default:
			msg += fmt.Sprintf("%s: %s歩\n", month, stepComma)
		}
	}
	msg += "\n"
	if yearMonth != "" {
		msg += yearMonth
	} else  {
		msg += "今月"
	}
	msg += fmt.Sprintf("の累計歩数は %s歩 ですニャー :cat2:", humanize.Comma(int64(steps.AllSteps)))
	return &slack.Msg{ResponseType: "in_channel", Text: msg}, nil
}

func getSteps(w http.ResponseWriter, req string) *stepsEachMonth {
	api := os.Getenv("BLOG_ACTIVITY_API")
	if api == "" {
		w.WriteHeader(http.StatusBadRequest)
		writeBody(w, []byte("必要な情報が足りていません"))
		return nil
	}

	sts := &stepsEachMonth{}
	client := resty.New()
	var resp *resty.Response
	var err error
	var path string
	if req != "" {
		path = fmt.Sprintf("%s/meru/steps/each_month/%s", api, req)
	} else {
		path = fmt.Sprintf("%s/meru/steps/each_month", api)
	}
	resp, err = client.R().EnableTrace().SetResult(sts).Get(path)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		writeBody(w, []byte(err.Error()))
		return nil
	}
	if resp == nil {
		w.WriteHeader(http.StatusInternalServerError)
		writeBody(w, []byte("予期せぬエラーが発生しています"))
		return nil
	}
	if resp.StatusCode() != http.StatusOK {
		w.WriteHeader(http.StatusInternalServerError)
		writeBody(w, []byte("歩数情報取得に失敗しました"))
		return nil
	}
	return sts
}
