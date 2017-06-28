package controllers

import (
	"blog/models"
)

type PostController struct {
	MyController
}

func (this *PostController) Get() {
	this.Data["Post"]= models.GetPosts()
	this.Layout = "index.html"
	this.TplName = "content.html"
}

//arama
func (this *PostController) Post() {
	txtAra :=this.GetString("ara")
	this.Data["Post"]= models.PostAra(txtAra)
	this.Data["txtAra"] =txtAra
	this.Layout = "index.html"
	this.TplName = "content.html"
}

func (this *PostController) Kategori() {
	id:= this.Ctx.Input.Param(":id")
	this.Data["Post"]= models.GetPostsInThisCatogory(id)
	this.Layout = "index.html"
	this.TplName = "content.html"
}

func (this *PostController) Detay() {
	id:= this.Ctx.Input.Param(":id")
	this.Data["Post"]=models.GetPost(id)
	this.Layout = "index.html"
	this.TplName = "detail.html"
}
