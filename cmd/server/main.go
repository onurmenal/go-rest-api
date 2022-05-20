package main

import (
	"context"
	"fmt"

	"github.com/onurmenal/go-rest-api/internal/comment"
	"github.com/onurmenal/go-rest-api/internal/db"
	transportHttp "github.com/onurmenal/go-rest-api/internal/transport/http"
)

func Run() error {
	fmt.Println("Starting up API")

	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to the database")
		return err
	}

	if err := db.Ping(context.Background()); err != nil {
		return err
	}

	if err := db.MigrateDB(); err != nil {
		fmt.Println("failed to migrate database")
		return err
	}

	cmtService := comment.NewService(db)

	httpHandler := transportHttp.NewHandler(cmtService)
	if err := httpHandler.Serve(); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
