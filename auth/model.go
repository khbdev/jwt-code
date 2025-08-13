package auth

type User struct {
    ID       int
    Email    string
    Password string
    Role string 
}


type LoginRequest struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

type LoginResponse struct {
    AccessToken string `json:"access_token"`
    RefreshToken string `json:refresh_token`
}