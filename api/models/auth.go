package models

type CustomerLoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}



type AuthInfo struct {
	UserID   string `json:"user_id"`
	UserRole string `json:"user_role"`
}

type UserLoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
type UserLoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
