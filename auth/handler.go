package auth

import (
	"encoding/json"
	

	"net/http"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
    var user User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, "invalid request", http.StatusBadRequest)
        return
    }

    if err := Register(user); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]string{
        "message": "registered successfully",
    })
}


func LoginHandler(w http.ResponseWriter, r *http.Request) {
    var creds LoginRequest
    if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
        http.Error(w, "invalid login", http.StatusBadRequest)
        return
    }

  accessToken, refreshToken, err := Login(creds.Email, creds.Password)
    if err != nil {
        http.Error(w, err.Error(), http.StatusUnauthorized)
        return
    }

    json.NewEncoder(w).Encode(LoginResponse{AccessToken: accessToken, RefreshToken: refreshToken})
}


func RefreshHandler(w http.ResponseWriter, r *http.Request) {
    var req struct {
        RefreshToken string `json:"refresh_token"`
    }

    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    newAccessToken, err := Refresh(req.RefreshToken)
    if err != nil {
        http.Error(w, "Invalid refresh token", http.StatusUnauthorized)
        return
    }

    json.NewEncoder(w).Encode(map[string]string{
        "access_token": newAccessToken,
    })
}

