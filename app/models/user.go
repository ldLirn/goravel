package models

import (
	"github.com/goravel/framework/database/orm"
	"github.com/goravel/framework/facades"
	"golang.org/x/crypto/bcrypt"
	"goravel/app/http/common"
	"log"
)

type User struct {
	orm.Model
	Name     string
	Password string
	orm.SoftDeletes
}

type ReturnUser struct {
	common.ReturnUser
}

//根据用户名查询
func (u *User) GetUser() (ReturnUser, error) {
	var ru ReturnUser
	err := facades.Orm.Query().Table("users").Where("name", u.Name).Find(&ru)

	if err != nil {
		return ru, err
	}
	if ru.Name == "" {
		return ru, err
	}
	return ru, err
}

//密码加密
func (u *User) BcPassword() (string, error) {
	//log.Println(u.Password)
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	return string(hash), err
}

//对比密码
func (u *User) ContrastPassword(pwd2 string) (ReturnUser, string) {
	data, _ := u.GetUser()
	var ru ReturnUser
	log.Println(data.Name)
	if data.Name == "" {
		return ReturnUser{}, "用户不存在"
	}

	err := bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(pwd2))
	log.Println(err)
	log.Println(pwd2)
	if err == nil {
		return data, ""
	} else {
		return ru, "密码错误"
	}
}
