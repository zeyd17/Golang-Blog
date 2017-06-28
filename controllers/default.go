package controllers

import (
	"github.com/astaxie/beego"
	"blog/models"
)

type Page struct {
	Title string
	Kategoriler []*models.Kategori
}

type MyController struct {
	beego.Controller
}

var sessionName = beego.AppConfig.String("SessionName")

func (this *MyController) Prepare() {
	kategoriler:=models.GetKategoriler()
	page :=Page{Title:"Golang ",Kategoriler:kategoriler}
	this.Data["page"] = &page
}

type KategoriController struct {
	MyController
}
