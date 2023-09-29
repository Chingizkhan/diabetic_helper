package diabet_helper

import "diabetHelperTelegramBot/bot/config"

const errPrefix = "grpcClient:bot->diabetHelper"

type ClientGrpc struct {
	addr string
}

func NewClient() ClientGrpc {
	return ClientGrpc{addr: config.Get().Service.DiabetHelper}
}
