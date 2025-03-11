package app

import (
	"context"
	"github.com/Rasikrr/learning_platform_core/application"
	authC "github.com/Rasikrr/learning_platform_gateway/internal/clients/auth"
	usersC "github.com/Rasikrr/learning_platform_gateway/internal/clients/users"
	"github.com/Rasikrr/learning_platform_gateway/internal/envs"
	"github.com/Rasikrr/learning_platform_gateway/internal/ports/http"
)

type App struct {
	*application.App
	authClient  authC.Client
	usersClient usersC.Client
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
	a.authClient, err = authC.NewClient(ctx, a.Config().Env.Get(envs.AuthGRPcAddress).GetString())
	if err != nil {
		return err
	}
	return nil
}

func (a *App) initServices(_ context.Context) error {
	return nil
}

func (a *App) initHTTP(_ context.Context) error {
	http.NewServer(
		a.HTTPServer(),
		a.usersClient,
		a.authClient,
	)
	return nil
}
