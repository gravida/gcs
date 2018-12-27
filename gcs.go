package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gravida/gcs/models"
	"github.com/gravida/gcs/pkg/settings"
	"github.com/gravida/gcs/routers"
	_ "github.com/mattn/go-sqlite3"
)

// GOOS=linux GOARCH=amd64 go build -o coolgo_linux *.go
// GOOS=windows GOARCH=amd64 go build -o coolgo_win *.go
func main() {

	settings.Setup()
	models.Setup()

	g := routers.InitRouter()
	g.Run(":54321")
}
