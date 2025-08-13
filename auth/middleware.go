package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"os"
	"strings"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			writeJSONError(w, "missing Authorization header", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			writeJSONError(w, "invalid Authorization format", http.StatusUnauthorized)
			return
		}
		token := parts[1]

		tokenParts := strings.Split(token, ".")
		if len(tokenParts) != 3 {
			writeJSONError(w, "invalid token structure", http.StatusUnauthorized)
			return
		}

		header := tokenParts[0]
		payload := tokenParts[1]
		signature := tokenParts[2]

		secret := os.Getenv("JWT_SECRET")
		if secret == "" {
			writeJSONError(w, "server misconfigured: missing JWT_SECRET", http.StatusInternalServerError)
			return
		}

		expectedSig := createHMAC(header+"."+payload, secret)

		if !hmac.Equal([]byte(signature), []byte(expectedSig)) {
			writeJSONError(w, "invalid token signature", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func createHMAC(data, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	return strings.TrimRight(base64.RawURLEncoding.EncodeToString(h.Sum(nil)), "=")
}

func writeJSONError(w http.ResponseWriter, message string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(ErrorResponse{Error: message})
}
