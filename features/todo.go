package features

import (
	"errors"
	"time"

	"github.com/gofrs/uuid"
	"gopkg.in/guregu/null.v4"
)

type todo struct {
	Id          uuid.UUID `db:"id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	IsCompleted bool      `db:"is_completed"`
	CreatedAt   time.Time `db:"created_at"`
	CreatedBy   string    `db:"created_by"`
	UpdatedAt   time.Time `db:"updated_at"`
	UpdatedBy   string    `db:"updated_by"`
	DeletedAt   null.Time `db:"deleted_at"`
	DeletedBy   string    `db:"deleted_by"`
}

func NewTodo(title string, description string) (todo, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return todo{}, err
	}
	if title == "" || len(title) > 100 {
		return todo{}, errors.New("title is empty or too long")
	}
	if len(description) > 500 {
		return todo{}, errors.New("description is too long")
	}

	return todo{
		Id:          id,
		Title:       title,
		Description: description,
		IsCompleted: false,
		CreatedAt:   time.Now(),
		CreatedBy:   "system",
		UpdatedAt:   time.Now(),
		UpdatedBy:   "system",
		DeletedAt:   null.Time{},
		DeletedBy:   "",
	}, nil
}
