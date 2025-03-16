package middlewares

import (
	"errors"
	"github.com/Rasikrr/learning_platform_core/api"
	"github.com/Rasikrr/learning_platform_core/http/session"
	usersC "github.com/Rasikrr/learning_platform_gateway/internal/clients/users"
	"log"
	"net/http"
)

type EnrollMiddleware struct {
	usersClient usersC.Client
}

func NewEnrollMiddleware(usersClient usersC.Client) *EnrollMiddleware {
	return &EnrollMiddleware{
		usersClient: usersClient,
	}
}

func (m *EnrollMiddleware) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("entering enroll middleware")
		ses, err := session.GetFromCtx(r.Context())
		if err != nil {
			api.SendError(w, http.StatusInternalServerError, err)
			return
		}
		courseID := r.URL.Query().Get("course_id")
		if courseID == "" {
			courseID = r.PathValue("course_id")
			if courseID == "" {
				api.SendError(w, http.StatusBadRequest, errors.New("course id is empty"))
				return
			}
		}
		enrolled, err := m.usersClient.CheckEnrollment(r.Context(), ses.UserID(), courseID)
		if err != nil {
			api.SendError(w, http.StatusInternalServerError, err)
			return
		}
		if !enrolled {
			api.SendError(w, http.StatusBadRequest, errors.New("user is not enrolled in course"))
			return
		}
		next.ServeHTTP(w, r)
	})
}
