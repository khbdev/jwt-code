package auth

import (
	"database/sql"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)



var DB *sql.DB


func Register(user User) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if  err != nil {
		return  err
	}
	_, err = DB.Exec("INSERT INTO users (email, password) VALUES (?, ?)", user.Email, string(hashed))
	return  err
}

func Login(email, password string) (string, error){
	var u User
	row := DB.QueryRow("SELECT id, email, password FROM users WHERE email = ?", email)
	err := row.Scan(&u.ID, &u.Email, &u.Password)
	 if err != nil {
        if err == sql.ErrNoRows {
            return "", errors.New("user not found")
        }
        return "", err
    }
	  if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
        return "", errors.New("invalid credentials")
    }
	 token, err := generateJWT(&u)
    if err != nil {
        return "", err
    }

    return token, nil
}

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