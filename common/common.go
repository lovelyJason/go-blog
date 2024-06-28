package common

import "go-blog/models"

var Template models.HtmlTemplate

func LoadTemplate() {
	Template = models.InitTemplate()
}
