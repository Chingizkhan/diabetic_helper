package start

import (
	"context"
	"diabetHelperTelegramBot/bot/service/grpc/diabet_helper"
	pb "diabetHelperTelegramBot/proto/diabetHelper"
	tele "gopkg.in/telebot.v3"
	"time"
)

func Handle(c tele.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	user := c.Sender()
	_, err := diabet_helper.NewClient().AddUser(ctx, &pb.AddUserRequest{
		Id:           user.ID,
		Username:     user.Username,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		LanguageCode: user.LanguageCode,
	})
	if err != nil {
		return err
	}

	return c.Send(`Вы успешно начали пользоваться ботом, нажмите /help для ознакомления с коммандами`)
}
