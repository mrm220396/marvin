package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

var (
	bot    = BotInfo{Version: 0.1, Name: "The Pub Waitress"}
	b, err = tb.NewBot(tb.Settings{
		// You can also set custom API URL.
		// If field is empty it equals to "https://api.telegram.org".

		Token:  os.Getenv("TELEGRAM_TOKEN"),
		Poller: &tb.LongPoller{Timeout: 1 * time.Second},
	})
	member      tb.ChatMember
	badwordlist = []string{"Fuck", "fuck", "shit", "mierda", "🖕"}
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Panic(port)
	}

	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle(tb.OnText, Handler)
	b.Handle("/info", ReturnInfo)
	b.Handle("/ban", Ban)
	b.Handle("/warn", Warn)
	b.Start()
}

func ReturnInfo(m *tb.Message) {
	text := fmt.Sprintf("My name is %s\nVersion %.2f", bot.Name, bot.Version)
	b.Send(m.Chat, text)
}

// Handler makes an evaluation of every message sent
func Handler(m *tb.Message) {

	fmt.Printf("%s, by %s %d\n", m.Text, m.Sender.FirstName, m.Sender.ID) // Terminal output

	// BadWordsParser treats every badword senteces
	BadwordsParser(m, m.Text)
}

// Warn to warn users
func Warn(m *tb.Message) {
	fmt.Println("inside warn func")
	member.User = m.ReplyTo.Sender
	user := GroupUser{User: member.User, Warnings: 1}
	result := fmt.Sprintf("Warned %s", user.User.Username)
	b.Send(m.Chat, result)
}

// Nan function bans members that are troublesome
func Ban(m *tb.Message) {
	// Finding the banned member you need
	// to reply him or her to be banned
	member.User = m.ReplyTo.Sender
	member.RestrictedUntil = 90
	b.Ban(m.Chat, &member)
}

func BadwordsParser(m *tb.Message, text string) {
	// badword var is a boolean which says if
	// there's a badword in the sentence or not
	var badword bool = false

	// result var will be formated the polite
	// version of every sentence which BadwrodsParser
	// function returns as true for badwords whithin it
	var result string

	//text = strings.ReplaceAll(text, ",", "")
	Newtext := strings.Split(text, " ")

	// For function will check every word sent
	// with this check will compare with the badword list
	for _, v := range Newtext {
		for _, word := range badwordlist {
			if strings.Contains(v, word) {
				text = strings.Replace(text, word, "🤬", 1)
				badword = true
			}
		}
	}
	if badword == true {
		result = fmt.Sprintf("Polite version of @%s 's message\n%s", m.Sender.Username, text)
		b.Delete(m)
		b.Send(m.Chat, result)
	}

}
