package auth

import (
	"database/sql"
	"errors"
	"golang.org/x/crypto/bcrypt"
)



var DB *sql.DB


func Register(user User) error {
	
	if user.Role == "" {
		user.Role = "user"
	}


	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = DB.Exec(
		"INSERT INTO users (email, password, role) VALUES (?, ?, ?)",
		user.Email, string(hashed), user.Role,
	)
	return err
}
func Login(email, password string) (string, string, error) {
    var u User
    row := DB.QueryRow("SELECT id, email, password, role FROM users WHERE email = ?", email)
    err := row.Scan(&u.ID, &u.Email, &u.Password, &u.Role)
    if err != nil {
        return "", "", errors.New("invalid credentials")
    }

    if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
        return "", "", errors.New("invalid credentials")
    }

    accessToken, err := generateJWT(&u)
    if err != nil {
        return "", "", err
    }

    refreshToken, err := GenerateRefreshToken(&u)
    if err != nil {
        return "", "", err
    }

    _, _ = DB.Exec("UPDATE users SET refresh_token = ? WHERE id = ?", refreshToken, u.ID)

    return accessToken, refreshToken, nil
}




