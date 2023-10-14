package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/jackc/pgx/v5/pgxpool"

	"todolist/features"
)

var getTodo *features.GetTodo
var updateTodo *features.UpdateTodo

func main() {
	ctx := context.Background()
	dbConfig := getConfig().DatabaseConfig

	dsn := fmt.Sprintf("postgresql://%v:%v@%v:%v/%v?sslmode=%v", dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name, dbConfig.SslMode)
	conn, err := pgxpool.New(ctx, dsn)
	if err != nil {
		panic(err)
	}

	getTodo = features.NewGetTodo(conn)
	updateTodo = features.NewUpdateTodo(conn)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/todos", func(r chi.Router) {
		r.Get("/", listTodos)
		r.With(middleware.BasicAuth("admin", map[string]string{"admin": "admin"})).Post("/", createTodo)
	})

	_ = http.ListenAndServe(":8080", r)
}

func listTodos(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	w.Header().Set("Content-Type", "application/json")
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))
	result := getTodo.GetList(ctx, offset)
	if getTodo.Err != nil {
		_ = render.Render(w, r, ErrRender(getTodo.Err))
		return
	}
	_ = json.NewEncoder(w).Encode(result)
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	w.Header().Set("Content-Type", "application/json")

	newTodo := &TodoRequest{}
	if err := render.Bind(r, newTodo); err != nil {
		_ = render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	updateTodo.CreateTodo(ctx, newTodo.Title, newTodo.Description)
	if updateTodo.Err != nil {
		_ = render.Render(w, r, ErrRender(updateTodo.Err))
		return
	}
	render.Status(r, http.StatusCreated)
}
