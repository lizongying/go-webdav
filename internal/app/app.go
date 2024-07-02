package app

import (
	"github.com/lizongying/go-webdav/internal/cli"
	"github.com/lizongying/go-webdav/internal/client"
	"github.com/lizongying/go-webdav/internal/config"
	"github.com/lizongying/go-webdav/internal/server"
	"go.uber.org/fx"
	"log"
)

type App struct {
}

func NewApp() (a *App) {
	a = new(App)
	return
}

func (a *App) Server() {
	constructors := []any{
		cli.NewCli,
		config.NewConfig,
		server.NewServer,
	}

	fx.New(
		fx.Provide(constructors...),
		fx.Invoke(func(server *server.Server, shutdowner fx.Shutdowner) {
			var err error
			if err = server.Run(); err != nil {
				log.Println(err)
				if err = shutdowner.Shutdown(); err != nil {
					log.Println(err)
				}
				return
			}

			return
		}),
	).Run()
}

func (a *App) Client() {
	constructors := []any{
		cli.NewCli,
		config.NewConfig,
		client.NewClient,
	}

	fx.New(
		fx.Provide(constructors...),
		fx.Invoke(func(client *client.Client, shutdowner fx.Shutdowner) {
			var err error
			if err = client.List(); err != nil {
				log.Println(err)
				if err = shutdowner.Shutdown(); err != nil {
					log.Println(err)
				}
				return
			}

			return
		}),
	).Run()
}
