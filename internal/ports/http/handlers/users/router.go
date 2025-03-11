package users

import (
	usersC "github.com/Rasikrr/learning_platform_gateway/internal/clients/users"
	"github.com/Rasikrr/learning_platform_gateway/internal/ports/http/middlewares"
	"github.com/go-chi/chi/v5"
)

type Controller struct {
	usersClient usersC.Client
	m           *middlewares.AuthMiddleware
}

func NewController(
	usersClient usersC.Client,
	m *middlewares.AuthMiddleware,
) *Controller {
	return &Controller{
		usersClient: usersClient,
		m:           m,
	}
}

func (c *Controller) Init(router *chi.Mux) {
	router.Route("/api/v1/users", func(r chi.Router) {
		r.With(c.m.Handle).Get("/me", c.getMyProfile)
		r.With(c.m.Handle).Put("/update", c.updateUser)
		r.With(c.m.Handle).Delete("/me/delete", c.deleteUser)

		r.Get("/{id}", c.getUser)
	})
}
