package middleware

import (
	"net/http"

	"github.com/joaodematejr/imersao22/go-gateway/internal/service"
)

type AuthMiddleware struct {
	accountService *service.AccountService
}

func NewAuthMiddleware(accountService *service.AccountService) *AuthMiddleware {
	return &AuthMiddleware{
		accountService: accountService,
	}
}

func (m *AuthMiddleware) Auth() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			apiKey := r.Header.Get("X-API-KEY")
			if apiKey == "" {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			_, err := m.accountService.FindByAPIKey(apiKey)
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
