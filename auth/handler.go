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
    w.Write([]byte("registered successfully"))
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
    var creds LoginRequest
    if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
        http.Error(w, "invalid login", http.StatusBadRequest)
        return
    }

    token, err := Login(creds.Email, creds.Password)
    if err != nil {
        http.Error(w, err.Error(), http.StatusUnauthorized)
        return
    }

    json.NewEncoder(w).Encode(LoginResponse{AccessToken: token})
}
