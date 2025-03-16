package auth

import (
	authC "github.com/Rasikrr/learning_platform_gateway/internal/clients/auth"
	"github.com/go-chi/chi/v5"
)

type Controller struct {
	authClient authC.Client
}

func NewController(authClient authC.Client) *Controller {
	return &Controller{
		authClient: authClient,
	}
}

func (c *Controller) Init(route *chi.Mux) {
	route.Route("/api/v1/auth", func(r chi.Router) {
		r.Post("/login", c.login)
		r.Post("/register", c.register)
		r.Post("/logout", c.logout)
		r.Post("/register/confirm", c.confirmRegister)
		r.Post("/refresh", c.refreshHandler)
		r.Post("/password/reset", c.resetPassword)
		r.Post("/password/reset/confirm", c.confirmResetPassword)
	})
}
