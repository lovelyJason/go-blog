package views

import (
	"go-blog/common"
	"go-blog/config"
	"net/http"
)

func (*HTMLApi) GetLogin(w http.ResponseWriter, r *http.Request) {
	login := common.Template.Login

	login.WriteData(w, config.Cfg.Viewer)
}
