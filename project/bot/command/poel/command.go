package poel

import (
	tele "gopkg.in/telebot.v3"
	"time"
)

func Handle(c tele.Context) error {
	err := c.Send("Напомню вам измерить уровень глюкозы через 2 часа!")
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
	time.Sleep(time.Hour * 2)
	err := c.Send("Время измерять глюкозу!🩸")
	if err != nil {
		return err
	}
	return nil
}
