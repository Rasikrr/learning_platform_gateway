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

func (c *Controller) Init(r *chi.Mux) {
	r.HandleFunc("POST /api/v1/auth/login", c.login)
	r.HandleFunc("POST /api/v1/auth/register", c.register)
	r.HandleFunc("POST /api/v1/auth/logout", c.logout)
	r.HandleFunc("POST /api/v1/auth/register/confirm", c.confirmRegister)
	r.HandleFunc("POST /api/v1/auth/refresh", c.refreshHandler)
	r.HandleFunc("POST /api/v1/auth/password/reset", c.resetPassword)
	r.HandleFunc("POST /api/v1/auth/password/reset/confirm", c.confirmResetPassword)
}
