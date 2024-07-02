package main

import (
	"github.com/lizongying/go-webdav/internal/app"
)

func main() {
	app.NewApp().Server()
}
