package http

import (
	"github.com/Rasikrr/learning_platform_core/http"
	authC "github.com/Rasikrr/learning_platform_gateway/internal/clients/auth"
	usersC "github.com/Rasikrr/learning_platform_gateway/internal/clients/users"
	"github.com/Rasikrr/learning_platform_gateway/internal/ports/http/handlers/auth"
	"github.com/Rasikrr/learning_platform_gateway/internal/ports/http/handlers/users"
	"github.com/Rasikrr/learning_platform_gateway/internal/ports/http/middlewares"
)

func NewServer(
	server *http.Server,
	usersClient usersC.Client,
	authClient authC.Client,
) {
	authMiddleware := middlewares.NewAuthMiddleware(authClient)

	server.WithControllers(
		users.NewController(usersClient, authMiddleware),
		auth.NewController(authClient),
	)
}
