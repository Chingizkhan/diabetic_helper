package main

import (
	"context"
	"diabetHelperTelegramBot/diabetHelper/bootstrap"
	"diabetHelperTelegramBot/diabetHelper/config"
	"diabetHelperTelegramBot/diabetHelper/server/grpc"
	"log"
	"os"
	"os/signal"
	"syscall"
)

const (
	envPath               = "./.env"
	postgresMigrationPath = "./storage/postgres/migration/"
)

func init() {
	bootstrap.Env(envPath)
	config.Init()
	bootstrap.PingDbConnect()
	bootstrap.Migrate(postgresMigrationPath)
	log.Println("init finished")
}

func main() {
	sigs := make(chan os.Signal, 1)
	defer close(sigs)

	done := make(chan struct{})
	defer close(done)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		grpc.GraceFulInit(ctx)
		done <- struct{}{}
	}()

	<-sigs
	cancel()
	<-done
}
