package middleware

import "net/http"

func (m *Middleware) authToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		_, err := m.Models.Tokens.AuthenticateToken(r)
		if err != nil {
			var payload struct {
				Error   bool   `json:"error"`
				Message string `json:"message"`
			}

			payload.Error = true
			payload.Message = "invalid authentication credentials"

			_ = m.App.WriteJSON(rw, http.StatusUnauthorized, payload)
		}
	})
}
