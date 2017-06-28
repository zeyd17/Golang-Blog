package Panel

import (
	"github.com/astaxie/beego"
	"blog/models"
)

type AdminController struct {
	beego.Controller
	User models.User
}

var sessionName = beego.AppConfig.String("SessionName")

func (this * AdminController) Prepare()  {

	sess := this.GetSession(sessionName)
	if sess != nil {
		m := sess.(map[string]models.User)
		this.User = m["user"]
	}else{
		this.Ctx.Redirect(302, "/")
	}
}

