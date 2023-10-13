package features

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UpdateTodo struct {
	pool *pgxpool.Pool
	Err  error
}

func NewUpdateTodo(writeDatabase *pgxpool.Pool) *UpdateTodo {
	return &UpdateTodo{
		pool: writeDatabase,
		Err:  nil,
	}
}

func (u *UpdateTodo) CreateTodo(ctx context.Context, title string, description string) {
	newTodo, err := NewTodo(title, description)
	u.Err = err
	_, u.Err = u.pool.Exec(ctx, `INSERT INTO todos (id, title, description, is_completed, created_at, created_by, updated_at, updated_by) 
VALUES ($1,$2,$3,$4,$5,$6,$7,$8)`,
		newTodo.Id, newTodo.Title, newTodo.Description, newTodo.IsCompleted, newTodo.CreatedAt, newTodo.CreatedBy, newTodo.UpdatedAt, newTodo.UpdatedBy)
}
