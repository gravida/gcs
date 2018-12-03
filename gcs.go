package main

import (
	"github.com/gravida/gcs/routers"
)

func main() {
	g := routers.InitRouter()
	g.Run()
}
