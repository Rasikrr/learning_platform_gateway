package commands

import (
	coursesC "github.com/Rasikrr/learning_platform_gateway/internal/clients/courses"
	"github.com/Rasikrr/learning_platform_gateway/internal/ports/http/middlewares"
	"github.com/go-chi/chi/v5"
)

type Controller struct {
	coursesClient     coursesC.Client
	submissionService submissionS.Service
	e                 *middlewares.EnrollMiddleware
	m                 *middlewares.AuthMiddleware
}

func NewController(
	coursesClient coursesC.Client,
	m *middlewares.AuthMiddleware,
	e *middlewares.EnrollMiddleware,
	submissionService submissionS.Service,
) *Controller {
	return &Controller{
		courseService:     courseService,
		submissionService: submissionService,
		m:                 m,
		e:                 e,
	}
}

func (c *Controller) Init(route *chi.Mux) {
	route.Route("/api/v1/courses", func(r chi.Router) {
		protected := r.With(c.m.Handle).With(c.e.Handle)

		protected.Post("/{course_id}/topic/{topic_id}/quiz/submit", c.submitQuiz)
		protected.Post("/{course_id}/topic/{topic_id}/task/{task_id}/submit", c.submitTask)
		protected.Post("/{course_id}/topic/{topic_id}/task/{task_id}/execute", c.executeTask)

		protected.Put("/{course_id}/topic/{topic_id}/quiz/reset", c.resetQuiz)

		protected.Delete("/{course_id}/topic/{topic_id}/task/{task_id}/reset", c.resetTask)
	})
}
