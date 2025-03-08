package http

import (
	"github.com/Rasikrr/learning_platform_core/http"
	"github.com/Rasikrr/learning_platform_gateway/internal/ports/http/handlers/users"
	usersS "github.com/Rasikrr/learning_platform_gateway/internal/services/users"
)

func NewServer(
	server *http.Server,
	usersService usersS.Service,
) {
	server.WithControllers(
		users.NewController(usersService),
	)
}
