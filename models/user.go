package models

import "time"

type User struct {
	Uid      string    `json:"uid"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Avatar   string    `json:"avatar"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}

type UserInfo struct {
	Uid      string `json:"uid"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
}
