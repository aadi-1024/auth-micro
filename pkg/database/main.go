package database

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"
	"time"
)

type Database struct {
	Pool *pgxpool.Pool
}

var dbTimeout = 3 * time.Second

var ctx = func() context.Context {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	return ctx
}

// InitDb assumes DSN is stored in an environment variable
func InitDb() (*Database, error) {
	dsn := os.Getenv("DSN")
	pool, err := pgxpool.New(ctx(), dsn)
	if err != nil {
		return nil, err
	}

	for i := 0; i < 5; i++ {
		log.Println("pinging database")
		err := pool.Ping(ctx())
		if err != nil {
			log.Println(err)
			continue
		}
		log.Println("successful")
		break
	}

	return &Database{
		pool,
	}, nil
}
