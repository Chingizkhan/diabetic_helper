package add

import (
	"context"
	"diabetHelperTelegramBot/bot/service/grpc/diabet_helper"
	pb "diabetHelperTelegramBot/proto/diabetHelper"
	"errors"
	tele "gopkg.in/telebot.v3"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func Handle(c tele.Context) error {
	sugarLevel := c.Text()

	if err := validate(&sugarLevel); err != nil {
		return c.Send("Incorrect format of sugar level - " + sugarLevel + ". Should input like 5.5")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	_, err := diabet_helper.NewClient().AddSL(ctx, &pb.AddSLRequest{
		Value:  sugarLevel,
		UserId: c.Sender().ID,
	})
	if err != nil {
		return c.Send("Can not save sugar level - " + sugarLevel + err.Error())
	}

	return c.Send("Saved sugar level - " + sugarLevel)
}

var reInt = regexp.MustCompile("(?i)[0-9]+\\.[0-9]+")

func validate(sugarLevel *string) error {
	*sugarLevel = strings.TrimSpace(*sugarLevel)
	*sugarLevel = strings.ReplaceAll(*sugarLevel, "-", "")
	*sugarLevel = strings.ReplaceAll(*sugarLevel, "+", "")
	*sugarLevel = strings.ReplaceAll(*sugarLevel, ",", ".")
	match := reInt.Match([]byte(*sugarLevel))

	if !match {
		return errors.New("unrecognised input")
	}

	_, err := strconv.ParseFloat(*sugarLevel, 64)
	if err != nil {
		return err
	}

	return nil
}
