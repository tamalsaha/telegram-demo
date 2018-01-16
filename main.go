package main

import (
	"fmt"
	"github.com/kimrgrey/go-telegram"
	"os"
	"net/url"
)

var client = telegram.NewClient(os.Getenv("BOT_KEY"))

func main() {
	me := client.GetMe()
	fmt.Println(me.UserName)

	params := url.Values{}
	params.Set("chat_id", "@mytest")
	params.Set("text", "Test")
	client.Call("sendMessage", params, nil)
}
