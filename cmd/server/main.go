package main

import (
	"context"
	"fmt"

	"github.com/onurmenal/go-rest-api/internal/comment"
	"github.com/onurmenal/go-rest-api/internal/db"
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
	fmt.Println(cmtService.GetComment(
		context.Background(),
		"904a8e95-b6a0-4f80-9142-5cf0a5895d25",
	))

	fmt.Println("Successfully connected and pinged database")

	return nil
}

func main() {
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
