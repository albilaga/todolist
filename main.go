package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	"todolist/features"
)

func main() {
	ctx := context.Background()
	dbConfig := getConfig().DatabaseConfig

	dsn := fmt.Sprintf("postgresql://%v:%v@%v:%v/%v?sslmode=%v", dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name, dbConfig.SslMode)
	conn, err := pgxpool.New(ctx, dsn)
	if err != nil {
		panic(err)
	}

	newUpdateTodo := features.NewUpdateTodo(conn)
	newUpdateTodo.CreateTodo(ctx, "testing", "this is for testing")
	if newUpdateTodo.Err != nil {
		panic(newUpdateTodo.Err)
	}
}
