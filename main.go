package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
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
	cfg := getConfig()
	dbConfig := cfg.DatabaseConfig
	appConfig := cfg.AppConfig

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	//serviceName := "todolist"
	//otelShutdown, err := setupOtelSDK(ctx, serviceName, "0.0.1")
	//if err != nil {
	//	return
	//}
	//defer func() {
	//	err = errors.Join(err, otelShutdown(context.Background()))
	//}()

	dsn := fmt.Sprintf("postgresql://%v:%v@%v:%v/%v?sslmode=%v", dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name, dbConfig.SslMode)
	conn, err := pgxpool.New(ctx, dsn)
	if err != nil {
		panic(err)
	}

	err = conn.Ping(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("Success ping database")

	getTodo = features.NewGetTodo(conn)
	updateTodo = features.NewUpdateTodo(conn)

	r := chi.NewRouter()
	//r.Use(middleware.Logger)
	//r.Use(otelchi.Middleware(serviceName, otelchi.WithChiRoutes(r)))

	r.Route("/todos", func(r chi.Router) {
		r.Get("/", listTodos)
		r.With(middleware.BasicAuth("admin", map[string]string{appConfig.Username: appConfig.Password})).Post("/", createTodo)
	})

	err = http.ListenAndServe(fmt.Sprintf("0.0.0.0:%v", appConfig.Port), r)
	if err != nil {
		panic(err)
	}
}

func listTodos(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	slog.Log(ctx, slog.LevelInfo, "get todos")
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
	slog.Log(ctx, slog.LevelInfo, "create todo")
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
	render.PlainText(w, r, "Ok")
}
