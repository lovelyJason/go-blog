package models

type LoginResponse struct {
	Token    string   `json:"token"`
	UserInfo UserInfo `json:"user_info"`
}
