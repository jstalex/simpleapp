package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"simpleapp/cmd/app/handler"
	"simpleapp/internal/config"
	"simpleapp/internal/service"
	"simpleapp/internal/storage"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := config.NewAppConfig()

	name := os.Getenv("PG_NAME")
	user := os.Getenv("PG_USER")
	pass := os.Getenv("PG_PASSWORD")
	host := os.Getenv("PG_HOST")

	url := fmt.Sprintf("postgres://%s:%s@%s:5432/%s", user, pass, host, name)

	conn, err := pgx.Connect(ctx, url)
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = conn.Ping(ctx)
	if err != nil {
		log.Printf(err.Error())
	}
	defer func() { _ = conn.Close(ctx) }()

	userStorage := storage.NewUserStorage(conn)
	appService := service.NewService(userStorage)
	appHandler := handler.NewHandler(appService)

	e := echo.New()

	e.Use(middleware.CORS())

	e.GET("/user", appHandler.GetUser)
	e.POST("/user", appHandler.PostUser)

	log.Fatal(e.Start(cfg.ListenAddr))
}
