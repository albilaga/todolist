package features

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type GetTodo struct {
	pool *pgxpool.Pool
	Err  error
}

const limit = 10

func NewGetTodo(readDatabase *pgxpool.Pool) *GetTodo {
	return &GetTodo{
		pool: readDatabase,
		Err:  nil,
	}
}

func (getTodo *GetTodo) GetList(ctx context.Context, offset int) {
	//getTodo.pool.Query()
}
