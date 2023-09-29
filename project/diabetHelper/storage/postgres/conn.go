package postgres

import (
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"log"
	"sync"
	"time"
)

var (
	onceDb sync.Once
	db     *sqlx.DB
)

const MaxOpenConns = 90
const MaxIdleConns = 90

var dbBuffConn = make(chan struct{}, 0.5*MaxOpenConns)

func HoldConn() {
	dbBuffConn <- struct{}{}
}

func GiveUpConn() {
	<-dbBuffConn
}

func GetDbConnection() *sqlx.DB {
	onceDb.Do(func() {
		db = getDbWithTicker()
		db.SetMaxOpenConns(MaxOpenConns)
		db.SetMaxIdleConns(MaxIdleConns)
		db.SetConnMaxIdleTime(20 * time.Minute)
	})
	return db
}

func getDbWithTicker() *sqlx.DB {
	done := make(chan struct{})
	var wg sync.WaitGroup

	wg.Add(1)
	go fatalEmptyChannelAfterTime(done, &wg, 60*time.Second)

	wg.Add(1)
	go pingDbInSomeTime(done, &wg, 5*time.Second)

	wg.Wait()
	return db
}

func pingDbInSomeTime(done chan struct{}, wg *sync.WaitGroup, duration time.Duration) {
	defer wg.Done()

	var err error
	db, err = getDb()
	if err == nil {
		done <- struct{}{}
		return
	}
	log.Printf("get db error=%v\n", err)

	ticker := time.NewTicker(duration)
	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			var err error
			db, err = getDb()
			if err != nil {
				log.Printf("get db error=%v\n", err)
				continue
			}
			done <- struct{}{}
			return
		}
	}
}

func fatalEmptyChannelAfterTime(done chan struct{}, wg *sync.WaitGroup, duration time.Duration) {
	defer wg.Done()
	tickerFail := time.NewTicker(duration)
	for {
		select {
		case <-done:
			return
		case <-tickerFail.C:
			log.Panicln("db init error")
			return
		}
	}
}

func getDb() (*sqlx.DB, error) {
	params := config.Get().Postgres

	var err error
	dsn := "postgres://" + params.User + ":" + params.Password + "@" + params.Host + ":" + params.Port + "/" + params.DbName + "?sslmode=disable"
	db, err := sqlx.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("can't open database: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("can't connect to database: %w", err)
	}

	return db, nil
}
