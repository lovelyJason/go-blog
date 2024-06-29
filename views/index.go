package views

import (
	"go-blog/common"
	"go-blog/service"
	"log"
	"net/http"
	"strconv"
)

// 给HTML这个变量添加了GetIndex方法
func (*HTMLApi) GetIndex(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/html")
	index := common.Template.Index
	// ！！获取请求参数
	if err := r.ParseForm(); err != nil {
		log.Println("parse form error", err)
		return
	}
	pageStr := r.Form.Get("page")
	page := 1
	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
	}
	pageSize := 10
	// 数据填充
	// 页面上涉及到的所有的数据，必须有定义
	homeData, err := service.GetAllInfo(page, pageSize)
	if err != nil {
		log.Printf("get all info failed, err:%v", err)
	}
	index.WriteData(w, homeData)

}
