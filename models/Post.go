package models

import "github.com/astaxie/beego/orm"

type Post struct {
	Id int
	Baslik string
	Ozet string
	Detay string
	Kategori string
}

func init()  {
	orm.RegisterModel(new(Post))
}

func GetPost(id string)  Post {
	var post Post
	err := orm.NewOrm().QueryTable("post").Filter("id",id).One(&post)
	if err != nil {
		return Post{}
	}
	return post
}

func GetPosts()  []*Post{
	var posts []*Post
	_, err := orm.NewOrm().QueryTable("post").OrderBy("-Id").All(&posts)
	if err != nil {
		return nil
	}
	return posts
}

func PostAra(txtAra string)  []*Post{
	var posts []*Post
	_, err := orm.NewOrm().QueryTable("post").Filter("Detay__icontains",txtAra).OrderBy("-Id").All(&posts)
	if err != nil {
		return nil
	}
	return posts
}

func GetPostsInThisCatogory(kategoriId string)  []*Post{
	var posts []*Post
	_, err := orm.NewOrm().QueryTable("post").Filter("Kategori",kategoriId).OrderBy("-id").All(&posts)
	if err != nil {
		return nil
	}
	return posts
}

func (this * Post) Sil() bool {
	o := orm.NewOrm()
	o.Using("default")
	_, err := o.Delete(this)
	if err ==nil {
		return  true
	}
	return false;
}

func (this * Post) Guncelle() bool  {
	o := orm.NewOrm()
	o.Using("default")
	_, err := o.Update(this)
	if err ==nil {
		return  true
	}
	return  false
}

func (this *Post) Kaydet() (int64, error)  {
	o := orm.NewOrm()
	o.Using("default")
	id, err := o.Insert(this)
	if err ==nil {
		return  id ,nil
	}else{
		return  -1 ,err
	}
}