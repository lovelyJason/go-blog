package api

import (
	"go-blog/service"
	"net/http"
)

func (*Api) Login(w http.ResponseWriter, r *http.Request) {
	// 注意：接收json
	//params := common.GetRequestJsonParams(r)
	//username := params["username"]
	//password := params["password"]

	//common.SuccessHandle(w, data)
	service.Login("jasonhuang", "123456")
}
