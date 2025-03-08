package users

import (
	usersS "github.com/Rasikrr/learning_platform_gateway/internal/services/users"
	"github.com/go-chi/chi/v5"
)

type Controller struct {
	usersService usersS.Service
}

func NewController(
	usersService usersS.Service,
) *Controller {
	return &Controller{
		usersService: usersService,
	}
}

func (c *Controller) Init(r *chi.Mux) {
	//r.HandleFunc("GET /api/v1/users/me", c.m.Handle(c.getMyProfile))
	r.HandleFunc("GET /api/v1/users/{email}", c.getUser)
	//r.HandleFunc("PUT /api/v1/users/update", c.m.Handle(c.updateUser))
	//r.HandleFunc("DELETE /api/v1/users/me/delete", c.m.Handle(c.deleteUser))
}
