package main

import (
	"fmt"

	"github.com/zhas-off/production-rest-api/internal/comment"
	"github.com/zhas-off/production-rest-api/internal/db"
	transportHttp "github.com/zhas-off/production-rest-api/internal/transport/http"
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

	cmtService := comment.NewService(db)

	httpHandler := transportHttp.NewHandler(cmtService)
	if err := httpHandler.Serve(); err != nil {
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
