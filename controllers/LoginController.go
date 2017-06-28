package controllers

import (
	"blog/models"
)

type LoginController struct {
	MyController
}



func (this *LoginController) Get() {
	this.TplName = "login.html"
}

func (this *LoginController) Post() {
	l :=models.LoginModel{}
	//l.Username =this.GetString("Username")
	//l.Password =this.GetString("Password")
	err :=this.ParseForm(&l);
	if err == nil{
		l.Password =models.Strtomd5(l.Password)
		user,ex :=l.Login();
		if ex == nil {
			m := make(map[string]models.User)
			m["user"] = user
			this.SetSession(sessionName, m)
			this.Redirect("/Panel",302)
		}

	}

	this.Redirect("/Login",302)
}

func (this *LoginController) LogOut() {
	this.DestroySession()
	this.Ctx.Redirect(302, "/")
}