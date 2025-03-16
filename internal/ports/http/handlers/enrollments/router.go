package enrollments

import (
	"github.com/Rasikrr/learning_platform_gateway/internal/ports/http/middlewares"
	"github.com/Rasikrr/learning_platform_gateway/internal/services/enrollments"
	"github.com/go-chi/chi/v5"
)

type Controller struct {
	enrollmentsService enrollments.Service
	m                  *middlewares.AuthMiddleware
}

func NewController(enrollmentsService enrollments.Service, m *middlewares.AuthMiddleware) *Controller {
	return &Controller{
		enrollmentsService: enrollmentsService,
		m:                  m,
	}
}

func (c *Controller) Init(route *chi.Mux) {
	route.Route("/api/v1/enrollments", func(r chi.Router) {
		protected := r.With(c.m.Handle)

		protected.Post("/enroll", c.enrollToCourse)
		protected.Get("/", c.getUserEnrollments)
	})

}
