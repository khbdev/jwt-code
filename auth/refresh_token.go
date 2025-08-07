package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateRefreshToken(user *User) (string, error) {
    secret := os.Getenv("REFRESH_SECRET")
    duration := 7 * 24 * time.Hour // 7 kun

    claims := jwt.MapClaims{
        "user_id": user.ID,
        "exp":     time.Now().Add(duration).Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(secret))
}