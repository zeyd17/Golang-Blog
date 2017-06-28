package routers

import (
	"blog/controllers"
	"github.com/astaxie/beego"
	"blog/controllers/Panel"
)

func init() {
	beego.Router("/", &controllers.PostController{})
	beego.Router("/Kategori/:id", &controllers.PostController{},"get:Kategori")
	beego.Router("/Detay/:id", &controllers.PostController{},"get:Detay")
	beego.Router("/Login",&controllers.LoginController{})
	beego.Router("/LogOut",&controllers.LoginController{},"get:LogOut")

	// panel.*  --> controller.Panel altındadır
	beego.Router("/Panel",&Panel.PanelController{})
	beego.AutoRouter(&Panel.PanelController{})
}
