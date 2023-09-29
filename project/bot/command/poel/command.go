package poel

import (
	tele "gopkg.in/telebot.v3"
	"time"
)

func Handle(c tele.Context) error {
	err := c.Send("Ok, I will remind you to check sugar level in 10 seconds!")
	if err != nil {
		return err
	}

	err = checkSugarLevel(c)
	if err != nil {
		return err
	}

	return nil
}

func checkSugarLevel(c tele.Context) error {
	time.Sleep(time.Second * 10)
	err := c.Send("It's time to check sugar!")
	if err != nil {
		return err
	}
	return nil
}
