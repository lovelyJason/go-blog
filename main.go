package main

import (
	"encoding/json"
	"go-blog/common"
	"go-blog/router"
	"log"
	"net/http"
)

type IndexData struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func getIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var indexData = IndexData{
		Title: "hello",
		Desc:  "world",
	}
	jsonStr, _ := json.Marshal(indexData)
	w.Write(jsonStr)
}

func init() {
	common.LoadTemplate()
}

func main() {
	log.Println("Server is starting on http://localhost:5174")

	server := http.Server{
		Addr: "127.0.0.1:5174",
	}
	router.Router()
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
