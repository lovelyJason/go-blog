package common

import (
	"encoding/json"
	"go-blog/config"
	"go-blog/models"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

// 这里封装的公共函数

var Template models.HtmlTemplate

func LoadTemplate() {
	w := sync.WaitGroup{}
	w.Add(1)
	// 注意：耗时操作，扔到协程里面
	go func() {
		var err error
		Template, err = models.InitTemplate(config.Cfg.System.CurrentDir + "/template/")
		if err != nil {
			panic(err)
		}
		w.Done()
	}()
}

func SuccessHandle(w http.ResponseWriter, data interface{}) {

	var result models.Result
	result.Code = 200
	result.Error = ""
	result.Data = data
	resultJson, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(resultJson)
	if err != nil {
		log.Println(err)
	}
}

// 获取json请求参数
func GetRequestJsonParams(r *http.Request) map[string]interface{} {
	var params map[string]interface{}
	body, _ := ioutil.ReadAll(r.Body) // 注意：这里即将废弃
	_ = json.Unmarshal(body, &params)
	return params
}
