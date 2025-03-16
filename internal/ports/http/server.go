package http

import (
	"github.com/Rasikrr/learning_platform_core/http"
	authC "github.com/Rasikrr/learning_platform_gateway/internal/clients/auth"
	coursesC "github.com/Rasikrr/learning_platform_gateway/internal/clients/courses"
	usersC "github.com/Rasikrr/learning_platform_gateway/internal/clients/users"
	"github.com/Rasikrr/learning_platform_gateway/internal/ports/http/handlers/auth"
	"github.com/Rasikrr/learning_platform_gateway/internal/ports/http/handlers/enrollments"
	enrollmentsS "github.com/Rasikrr/learning_platform_gateway/internal/services/enrollments"

	//coursesCommands "github.com/Rasikrr/learning_platform_gateway/internal/ports/http/handlers/courses/commands"
	coursesQueries "github.com/Rasikrr/learning_platform_gateway/internal/ports/http/handlers/courses/queries"
	"github.com/Rasikrr/learning_platform_gateway/internal/ports/http/handlers/users"
	"github.com/Rasikrr/learning_platform_gateway/internal/ports/http/middlewares"
)

func NewServer(
	server *http.Server,
	coursesClient coursesC.Client,
	usersClient usersC.Client,
	authClient authC.Client,
	enrollmentsService enrollmentsS.Service,
) {
	authMiddleware := middlewares.NewAuthMiddleware(authClient)
	enrollMiddleware := middlewares.NewEnrollMiddleware(usersClient)

	server.WithControllers(
		users.NewController(usersClient, authMiddleware),
		auth.NewController(authClient),
		coursesQueries.NewController(coursesClient, authMiddleware, enrollMiddleware),
		//coursesCommands.NewController(coursesClient, authMiddleware),
		enrollments.NewController(enrollmentsService, authMiddleware),
	)

}
