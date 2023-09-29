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

	wg := sync.WaitGroup{}
	var n = 5000
	wg.Add(n)
	for i := 0; i < n; i++ {
		go find(ctx, client, &wg)
	}
	wg.Wait()
	log.Println("index is", index)
}

var index int64

func find(ctx context.Context, client diabetHelper.DiabetHelperClient, wg *sync.WaitGroup) *diabetHelper.SugarLevels {
	defer wg.Done()
	point, err := client.FindSL(ctx, &diabetHelper.FindSLRequest{
		Pagination: &diabetHelper.Pagination{
			Page:   0,
			Limit:  10,
			Offset: 0,
		},
		Uuid:            "",
		Value:           "",
		Username:        "",
		CreatedAtStart:  0,
		CreatedAtFinish: 0,
	})
	if err != nil {
		log.Fatalln("error", err)
	}
	atomic.AddInt64(&index, 1)
	//log.Println(point)
	return point
}
