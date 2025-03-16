package queries

import (
	coursesC "github.com/Rasikrr/learning_platform_gateway/internal/clients/courses"
	"github.com/Rasikrr/learning_platform_gateway/internal/ports/http/middlewares"
	"github.com/go-chi/chi/v5"
)

type Controller struct {
	coursesClient coursesC.Client
	e             *middlewares.EnrollMiddleware
	m             *middlewares.AuthMiddleware
}

func NewController(
	coursesClient coursesC.Client,
	m *middlewares.AuthMiddleware,
	e *middlewares.EnrollMiddleware,
) *Controller {
	return &Controller{
		coursesClient: coursesClient,
		m:             m,
		e:             e,
	}
}

func (c *Controller) Init(route *chi.Mux) {
	route.Route("/api/v1/courses", func(r chi.Router) {
		protected := r.With(c.m.Handle).With(c.e.Handle)

		r.Post("/", c.getCourses)
		r.Get("/categories", c.getCategories)
		r.Get("/{id}", c.getCourse)

		protected.Get("/{course_id}/topic/{topic_id}/content", c.getCourseTopicContent)
		protected.Get("/{course_id}/topic/{topic_id}/quizzes", c.getCourseTopicQuizzes)
		protected.Get("/{course_id}/topic/{topic_id}/tasks/{order}", c.getCourseTopicTasks)
	})
}
