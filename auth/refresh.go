package auth

import (
	"errors"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func Refresh(refreshToken string) (string, error) {
    secret := os.Getenv("REFRESH_SECRET")

    token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
        return []byte(secret), nil
    })

    if err != nil || !token.Valid {
        return "", errors.New("invalid refresh token")
    }

    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok || claims["user_id"] == nil {
        return "", errors.New("invalid token claims")
    }

    // Extract user ID from claims
    userID := int(claims["user_id"].(float64))

    // (Ixtiyoriy) DBdan userni topish
    var user User
    err = DB.QueryRow("SELECT id, email FROM users WHERE id = ?", userID).Scan(&user.ID, &user.Email)
    if err != nil {
        return "", errors.New("user not found")
    }

    // Yangi access token berish
    return generateJWT(&user)
}
