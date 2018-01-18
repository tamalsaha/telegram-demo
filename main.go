package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/golang/glog"
	"github.com/tamalsaha/go-oneliners"
)

type ErrorResponse struct {
	Ok          bool   `json:"ok"`
	ErrorCode   int    `json:"error_code"`
	Description string `json:"description"`
}

func main() {
	token := os.Getenv("BOT_KEY")
	channel := "@hoannv_alert"

	u := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", token)

	oneliners.FILE(u)

	data := url.Values{}
	data.Set("chat_id", channel)
	data.Set("text", "this is from GO")
	resp, err := http.PostForm(u, data)
	if err != nil {
		glog.Fatalln(err)
	}
	if resp.StatusCode != http.StatusOK {
		var r ErrorResponse
		err := json.NewDecoder(resp.Body).Decode(&r)
		if err == nil && !r.Ok {
			glog.Warningf("failed to send message to channel %s. Reason: %d - %s", channel, r.ErrorCode, r.Description)
		}
	}
	resp.Body.Close()
}
