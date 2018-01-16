package main

import (
	"fmt"
	"github.com/kimrgrey/go-telegram"
	"os"
)

var client = telegram.NewClient(os.Getenv("BOT_KEY"))

func main() {
	me := client.GetMe()
	fmt.Println(me.UserName)
}
