package main

import (
	"context"
	"github.com/Rasikrr/learning_platform_gateway/internal/app"
	"log"
)

const (
	appName = "gateway"
)

func main() {
	ctx := context.Background()
	app, err := app.NewApp(ctx, appName)
	if err != nil {
		panic(err)
	}
	if err := app.Start(ctx); err != nil {
		log.Println(err)
	}
}
