package views

import (
	"go-blog/config"
	"go-blog/models"
	"html/template"
	"log"
	"net/http"
	"time"
)

func IsODD(num int) bool {
	return num%2 == 0
}

func GetNextName(strs []string, index int) string {
	return strs[index+1]
}

func Date(format string) string {
	return time.Now().Format(format)
}

// 给HTML这个变量添加了GetIndex方法
func (*HTMLApi) GetIndex(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/html")
	t := template.New("index.html")
	path := config.Cfg.System.CurrentDir

	// 解析各个子页面
	home := path + "/template/home.html"
	header := path + "/template/layout/header.html"
	footer := path + "/template/layout/footer.html"
	personal := path + "/template/layout/personal.html"
	postList := path + "/template/layout/post-list.html"
	pagination := path + "/template/layout/pagination.html"
	// 给模板注入方法
	t.Funcs(template.FuncMap{"isODD": IsODD, "getNextName": GetNextName, "date": Date})
	t, err := t.ParseFiles(path+"/template/index.html", home, header, footer, personal, postList, pagination)

	if err != nil {
		log.Println("出错了", err)
	}

	// 数据填充
	// 页面上涉及到的所有的数据，必须有定义
	var categorys = []models.Category{
		{
			Cid:  1,
			Name: "go",
		},
	}
	var posts = []models.PostMore{
		{
			Pid:          1,
			Title:        "go博客",
			Content:      "内容",
			UserName:     "jasonhuang",
			ViewCount:    123,
			CreateAt:     "2022-02-20",
			CategoryId:   1,
			CategoryName: "go",
			Type:         0,
		},
	}
	var homeData = &models.HomeData{
		config.Cfg.Viewer,
		categorys,
		posts,
		1,
		1,
		[]int{1},
		true,
	}

	t.Execute(w, homeData)
}
