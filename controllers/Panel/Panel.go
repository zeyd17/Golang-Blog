package Panel

import (
	"blog/models"
	"strconv"
)

type PanelController struct {
	AdminController
}

func (this *PanelController) Get()  {

	this.Data["User"] =models.GetUsers()
	this.Layout = "panel/index.html"
	this.TplName = "panel/kullanici.html"
}

func (this *PanelController) Post()  {
	if this.Ctx.Input.IsPost() {
		kullanici := models.User{}
		err :=this.ParseForm(&kullanici);
		if err == nil {
			kullanici.Pass =models.Strtomd5(kullanici.Pass)
			kullanici.Kaydet()
		}
	}
	this.Redirect("/Panel",302)
}

func (this * PanelController) KullaniciGuncelle()  {
	maps :=this.Ctx.Input.Params()
	id,_:=strconv.Atoi(maps["0"])
	this.Data["kullanici"] =models.GetUser(id)
	if this.Ctx.Input.IsPost() {
		kullanici := models.User{}
		err :=this.ParseForm(&kullanici);
		if err == nil {
			kullanici.Pass =models.Strtomd5(kullanici.Pass)
			kullanici.Guncelle()
		}
		this.Redirect("/Panel",302)
	}
	this.Data["User"] =models.GetUsers()
	this.Layout = "panel/index.html"
	this.TplName = "panel/kullanici.html"
}

func (this *PanelController) KullaniciSil() {
	maps :=this.Ctx.Input.Params()
	id,_:=strconv.Atoi(maps["0"])
	kullanici :=models.User{Id:id}
	kullanici.Sil()
	this.Redirect("/Panel",302)
}

/**************** Kategori **************/
func (this * PanelController) Kategori()  {

	if this.Ctx.Input.IsPost() {
		kategori := models.Kategori{}
		err :=this.ParseForm(&kategori);
		if err == nil {
			kategori.Kaydet()
		}
	}
	this.Data["Kategoriler"] =models.GetKategoriler()
	this.Layout = "panel/index.html"
	this.TplName = "panel/kategori.html"
}

func (this * PanelController) KategoriDuzenle()  {

	maps :=this.Ctx.Input.Params()
	id,_:=strconv.Atoi(maps["0"])
	if this.Ctx.Input.IsPost() {
		kategori := models.Kategori{}
		err :=this.ParseForm(&kategori);
		if err == nil {
			//beego.Debug("Data geldi",kategori)
			kategori.Guncelle()
		}
		this.Redirect("/Panel/Kategori",302)
	}

	this.Data["Kategori"] =models.GetKategori(id)
	this.Data["Kategoriler"] =models.GetKategoriler()
	this.Layout = "panel/index.html"
	this.TplName = "panel/kategori.html"
}

func (this * PanelController) KategoriSil()  {
	maps :=this.Ctx.Input.Params()
	id,_:=strconv.Atoi(maps["0"])
	kategori := models.Kategori{Id:int64(id)}
	kategori.Sil()
	this.Redirect("/Panel/Kategori",302)
}


/****************Post ****************/
func (this * PanelController) Posts()  {
	this.Data["Post"] =models.GetPosts()
	this.Layout = "panel/index.html"
	this.TplName = "panel/post.html"
}

func (this * PanelController) PostSil()  {
	maps :=this.Ctx.Input.Params()
	id,_:=strconv.Atoi(maps["0"])
	post := models.Post{Id:int(id)}
	post.Sil()
	this.Redirect("/Panel/Posts",302)
}

func (this * PanelController) PostEkle(){
	if this.Ctx.Input.IsPost() {
		post := models.Post{}
		err :=this.ParseForm(&post);
		if err == nil {
			post.Kaydet()
		}
		this.Redirect("/Panel/PostEkle",302)
	}
	this.Data["Kategoriler"] =models.GetKategoriler()
	this.Layout = "panel/index.html"
	this.TplName = "panel/postEkle.html"
}

func (this * PanelController) PostGuncelle(){
	maps :=this.Ctx.Input.Params()
	id,_:=maps["0"]
	if this.Ctx.Input.IsPost() {
		post := models.Post{}
		err :=this.ParseForm(&post);
		if err == nil {
			post.Guncelle()
		}
		this.Redirect("/Panel/Posts",302)
	}
	this.Data["Post"] =models.GetPost(id)
	this.Data["Kategoriler"] =models.GetKategoriler()
	this.Layout = "panel/index.html"
	this.TplName = "panel/postEkle.html"
}