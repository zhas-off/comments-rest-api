package main

import (
	"fmt"

	"github.com/zhas-off/production-rest-api/internal/db"
)

func Run() error {
	fmt.Println("Starting our application")

	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to the database")
		return err
	}
	if err := db.MigrateDB(); err != nil {
		fmt.Println("failed to migrate database")
		return err
	}
	return nil
}

func main() {
	fmt.Println("Hello rest api")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
