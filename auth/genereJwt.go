package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)



func generateJWT(user *User) (string, error) {
    secret := os.Getenv("JWT_SECRET")
    expireMin, _ := time.ParseDuration(os.Getenv("JWT_EXPIRE_MINUTES") + "m")

    claims := jwt.MapClaims{
        "user_id": user.ID,
        "email":   user.Email,
        "exp":     time.Now().Add(expireMin).Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(secret))
}
