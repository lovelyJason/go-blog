package router

import (
	"go-blog/api"
	"go-blog/views"
	"net/http"
)

func Router() {
	http.HandleFunc("/", views.HTML.GetIndex)
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))
	//http.HandleFunc("/test.html", getIndexHtml)
	http.HandleFunc("/api/v1/post", api.API.UpdatePost)
}
