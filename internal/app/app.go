package app

import (
	"context"
	"github.com/Rasikrr/learning_platform_core/application"
	usersC "github.com/Rasikrr/learning_platform_gateway/internal/clients/users"
	"github.com/Rasikrr/learning_platform_gateway/internal/envs"
	"github.com/Rasikrr/learning_platform_gateway/internal/ports/http"
	usersS "github.com/Rasikrr/learning_platform_gateway/internal/services/users"
)

type App struct {
	*application.App
	usersService usersS.Service
	usersClient  usersC.Client
}

func NewApp(ctx context.Context, name string) (*App, error) {
	app := &App{
		App: application.NewApp(ctx, name),
	}
	if err := app.Init(ctx); err != nil {
		return nil, err
	}

	return app, nil
}

func (a *App) Init(ctx context.Context) error {
	for _, init := range []func(context.Context) error{
		a.initClients,
		a.initServices,
		a.initHTTP,
	} {
		if err := init(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (a *App) initClients(ctx context.Context) error {
	var err error
	a.usersClient, err = usersC.NewClient(ctx, a.Config().Env.Get(envs.UsersGRPcAddress).GetString())
	if err != nil {
		return err
	}
	return nil
}

func (a *App) initServices(ctx context.Context) error {
	a.usersService = usersS.NewService(a.usersClient)
	return nil
}

func (a *App) initHTTP(_ context.Context) error {
	http.NewServer(
		a.HTTPServer(),
		a.usersService,
	)
	return nil
}
