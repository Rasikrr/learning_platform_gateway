package middlewares

import (
	"context"
	"errors"
	"github.com/Rasikrr/learning_platform_core/api"
	"github.com/Rasikrr/learning_platform_core/http/session"
	authC "github.com/Rasikrr/learning_platform_gateway/internal/clients/auth"
	"log"
	"net/http"
)

const (
	authHeader = "Authorization"
)

type AuthMiddleware struct {
	authClient authC.Client
}

func NewAuthMiddleware(authClient authC.Client) *AuthMiddleware {
	return &AuthMiddleware{
		authClient: authClient,
	}
}

func (m *AuthMiddleware) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("entering auth middleware")
		ses, err := m.parseAuth(r)
		if err != nil {
			api.SendError(w, http.StatusUnauthorized, err)
			return
		}
		r = r.WithContext(context.WithValue(r.Context(), session.SessionKey, ses))
		next.ServeHTTP(w, r)
	})
}

func (m *AuthMiddleware) parseAuth(r *http.Request) (*session.Session, error) {
	token := r.Header.Get(authHeader)
	if token == "" {
		return nil, errors.New("authorization header is empty")
	}
	ctx := r.Context()
	ses, err := m.authClient.CheckToken(ctx, token)
	if err != nil {
		return nil, err
	}
	return ses, nil
}
