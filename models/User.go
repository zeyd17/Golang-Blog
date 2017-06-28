package models

import (
	"github.com/astaxie/beego/orm"
	"crypto/md5"
	"encoding/hex"
)

type User struct {
	Id int
	Name string
	Mail string
	Pass string
}

type LoginModel struct {
	Username string
	Password string `orm:"column(Pass)"`
}

func init() {
	orm.RegisterModel(new(User))
}

func Strtomd5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	rs := hex.EncodeToString(h.Sum(nil))
	return rs
}

func GetUser(id int)  User {
	var user User
	err :=orm.NewOrm().QueryTable("user").Filter("id",id).One(&user)
	if err != nil{
		return User{}
	}
	return  user
}

func GetUsers()  []*User{
	var users []*User
	_,err :=orm.NewOrm().QueryTable("user").All(&users)
	if err != nil{
		return nil
	}
	return  users
}

func (this * User) Sil() bool {
	o := orm.NewOrm()
	o.Using("default")
	_, err := o.Delete(this)
	if err ==nil {
		return  true
	}
	return false;
}

func (this * User) Guncelle() bool  {
	o := orm.NewOrm()
	o.Using("default")
	_, err := o.Update(this)
	if err ==nil {
		return  true
	}
	return  false
}

func (this *User) Kaydet() (int64, error)  {
	o := orm.NewOrm()
	o.Using("default")
	id, err := o.Insert(this)
	if err ==nil {
		return  id ,nil
	}else{
		return  -1 ,err
	}
}

func (this *LoginModel) Login()  (User,error){
	var user User
	cond := orm.NewCondition()

	cond1 := cond.And("Mail", this.Username).And("Pass", this.Password)

	err :=orm.NewOrm().QueryTable("user").SetCond(cond1).One(&user)

	if err != nil {
		return  User{} ,err
	}
	return user,nil
}

