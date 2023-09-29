package main

import (
	"diabetHelperTelegramBot/bot/bootstrap"
	"diabetHelperTelegramBot/bot/command/add"
	"diabetHelperTelegramBot/bot/command/poel"
	"diabetHelperTelegramBot/bot/command/statistic"
	"diabetHelperTelegramBot/bot/config"
	tele "gopkg.in/telebot.v3"
	"log"
	"time"
)

const (
	envPath = "./.env"
)

func init() {
	bootstrap.Env(envPath)
	config.Init()
}

type User struct {
	name string
}

func New(name string) User {
	return User{name: name}
}

func (u *User) Recipient() string {
	return u.name
}

const (
	me     = "395096494"
	alihan = "1321034160"
	some   = ""
)

func main() {
	pref := tele.Settings{
		Token:  config.Get().Token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/calc", func(c tele.Context) error {
		return c.Send("Hello!")
	})

	b.Handle("/poel", func(c tele.Context) error {
		log.Printf("%#v", c.Sender())
		return poel.Handle(c)
	})

	b.Handle("/stat", func(c tele.Context) error {
		return statistic.Handle(c)
	})

	//send, err := b.Send(&User{name: alihan}, "Ловим приход")
	//if err != nil {
	//	log.Println("err", err)
	//}
	//log.Println("send", send)

	b.Handle(tele.OnText, func(c tele.Context) error {
		// All the text messages that weren't
		// captured by existing handlers.

		//var (
		//	user = c.Sender()
		//	text = c.Text()
		//)
		//
		//// Instead, prefer a context short-hand:
		//return c.Send(fmt.Sprintf("Sorry @%s. I can't understand '%s'", user.Username, text))
		return add.Handle(c)
	})

	b.Start()
}
