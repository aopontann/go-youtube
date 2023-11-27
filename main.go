package main

import (
	"context"
	"fmt"
	"os"
	"net/http"
    
    "github.com/labstack/echo/v4"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
	"github.com/jackc/pgx/v5"
)

func main() {
	ctx := context.Background()
	s, err := youtube.NewService(ctx, option.WithAPIKey(os.Getenv("YOUTUBE_API_KEY")))

	conn, err := pgx.Connect(context.Background(), os.Getenv("POSTGRES_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	e := echo.New()
    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, World!")
    })

}
