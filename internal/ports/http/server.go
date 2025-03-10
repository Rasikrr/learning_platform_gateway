package http

import (
	"github.com/Rasikrr/learning_platform_core/http"
	authC "github.com/Rasikrr/learning_platform_gateway/internal/clients/auth"
	"github.com/Rasikrr/learning_platform_gateway/internal/ports/http/handlers/auth"
	"github.com/Rasikrr/learning_platform_gateway/internal/ports/http/handlers/users"
	usersS "github.com/Rasikrr/learning_platform_gateway/internal/services/users"
)

func NewServer(
	server *http.Server,
	usersService usersS.Service,
	authClient authC.Client,
) {
	server.WithControllers(
		users.NewController(usersService),
		auth.NewController(authClient),
	)
}
