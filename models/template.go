package models

import (
	"html/template"
	"io"
	"log"
	"time"
)

// 注意：这里的字段有Template
type TemplateBlog struct {
	*template.Template // template.New返回的类型
}

type HtmlTemplate struct {
	Index      TemplateBlog
	Category   TemplateBlog
	Custom     TemplateBlog
	Detail     TemplateBlog
	Login      TemplateBlog
	Pigeonhole TemplateBlog
	Writing    TemplateBlog
}

func IsODD(num int) bool {
	return num%2 == 0
}

func GetNextName(strs []string, index int) string {
	return strs[index+1]
}

func Date(format string) string {
	return time.Now().Format(format)
}

func DateDay(date time.Time) string {
	return date.Format("2006-01-02 15:04:05")
}

func (t *TemplateBlog) WriteData(w io.Writer, data interface{}) {
	errOuter := t.Execute(w, data)
	if errOuter != nil {
		_, err := w.Write([]byte(errOuter.Error()))
		if err != nil {
			return
		}
	}
}

func InitTemplate(templateDir string) (HtmlTemplate, error) {

	tp, err := readTemplate(
		[]string{"index", "category", "custom", "detail", "login", "pigeonhole", "writing"},
		templateDir,
	)

	var htmlTemplate HtmlTemplate

	if err != nil {
		return htmlTemplate, err
	}

	htmlTemplate.Index = tp[0]
	htmlTemplate.Category = tp[1]
	htmlTemplate.Custom = tp[2]
	htmlTemplate.Detail = tp[3]
	htmlTemplate.Login = tp[4]
	htmlTemplate.Pigeonhole = tp[5]
	htmlTemplate.Writing = tp[6]
	return htmlTemplate, nil
}

func readTemplate(templates []string, templateDir string) ([]TemplateBlog, error) {
	var templateBlogs []TemplateBlog
	for _, view := range templates {
		viewName := view + ".html"
		t := template.New(viewName)
		//path := config.Cfg.System.CurrentDir

		// 解析各个子页面
		home := templateDir + "home.html"
		header := templateDir + "layout/header.html"
		footer := templateDir + "layout/footer.html"
		personal := templateDir + "layout/personal.html"
		postList := templateDir + "layout/post-list.html"
		pagination := templateDir + "layout/pagination.html"
		// 给模板注入方法
		t.Funcs(template.FuncMap{"isODD": IsODD, "getNextName": GetNextName, "date": Date, "dateDay": DateDay})
		t, err := t.ParseFiles(templateDir+viewName, home, header, footer, personal, postList, pagination)

		if err != nil {
			log.Println("出错了", err)
			return nil, err
		}
		var tb TemplateBlog
		tb.Template = t
		templateBlogs = append(templateBlogs, tb) // 注意这里不是t
	}
	return templateBlogs, nil
}
