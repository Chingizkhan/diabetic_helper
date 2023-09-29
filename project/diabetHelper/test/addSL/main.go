package main

import (
	"context"
	"diabetHelperTelegramBot/diabetHelper/bootstrap"
	"diabetHelperTelegramBot/diabetHelper/config"
	"diabetHelperTelegramBot/proto/diabetHelper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"sync"
	"sync/atomic"
	"time"
)

func init() {
	bootstrap.Env("./.env")
	config.Init()
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	port := "localhost:" + config.Get().Port

	conn, err := grpc.DialContext(ctx, port, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	defer func() {
		if conn != nil {
			_ = conn.Close()
		}
	}()
	if err != nil {
		log.Println("Can not connect to", port, err)
		return
	}

	client := diabetHelper.NewDiabetHelperClient(conn)

	var n = 10
	wg := sync.WaitGroup{}

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			_, err := client.AddSL(ctx, &diabetHelper.AddSLRequest{
				Value:    "12",
				Username: "user",
			})
			if err != nil {
				log.Fatalln("AddSL error:", err)
				return
			}
			atomic.AddInt64(&index, 1)
		}(&wg)
	}
	wg.Wait()

	log.Println("index is -", index)
}

var index int64
