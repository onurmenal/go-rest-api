package main

import (
	"context"
	"fmt"

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

	fmt.Println("Successfully connected and pinged database")

	return nil
}

func main() {
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
