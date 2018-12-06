package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gravida/gcs/models"
	"github.com/gravida/gcs/pkg/settings"
	"github.com/gravida/gcs/routers"
)

func main() {

	settings.Setup()
	models.Setup()

	g := routers.InitRouter()
	g.Run(":54321")
}
