package router

import (
	"go-blog/api"
	"go-blog/views"
	"net/http"
)

func Router() {
	// 注意：优化为上下文更好
	http.HandleFunc("/", views.HTML.GetIndex)
	http.HandleFunc("/login", views.HTML.GetLogin)
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))
	http.HandleFunc("/api/v1/post", api.API.UpdatePost)
	http.HandleFunc("/api/v1/login", api.API.Login)
}
