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

func (g *GetTodo) GetList(ctx context.Context, offset int) []TodoList {
	items := make([]TodoList, 0)
	rows, err := g.pool.Query(ctx, `SELECT id,title,description,is_completed FROM todos LIMIT $1 OFFSET $2`, limit, offset)
	if err != nil {
		g.Err = err
		return items
	}
	defer rows.Close()

	for rows.Next() {
		var item TodoList
		g.Err = rows.Scan(&item.Id, &item.Title, &item.Description, &item.IsCompleted)
		items = append(items, item)
	}
	g.Err = rows.Err()
	return items
}
