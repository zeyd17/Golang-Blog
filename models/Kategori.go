package models

import (
	"github.com/astaxie/beego/orm"
)

type Kategori struct {
	Id int64
	Kategori string
}

func init()  {
	orm.RegisterModel(new(Kategori))
}

func GetKategori(id int) Kategori {
	var kategori Kategori
	err:=orm.NewOrm().QueryTable("Kategori").Filter("id",id).One(&kategori)
	if err != nil {
		return  Kategori{}
	}
	return  kategori
}

func GetKategoriler() []*Kategori {
	var kategoriler []*Kategori
	_,err :=orm.NewOrm().QueryTable("Kategori").All(&kategoriler)
	if err !=nil {
		return  nil
	}
	return kategoriler
}

func  (this *Kategori) Sil() (bool) {
	o := orm.NewOrm()
	o.Using("default")
	_, err := o.Delete(this)
	if err == nil {
		return  true
	}
	return false;
}

func (this * Kategori) Guncelle() bool  {
	o := orm.NewOrm()
	o.Using("default")
	_, err := o.Update(this)
	if err ==nil {
		return  true
	}
	return  false
}

func (this *Kategori) Kaydet() (int64, error)  {
	o := orm.NewOrm()
	o.Using("default")
	id, err := o.Insert(this)
	if err ==nil {
		return  id ,nil
	}else{
		return  -1 ,err
	}
}