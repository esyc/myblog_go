package models

import (
	"crcblog/utils"
	"encoding/hex"
	"fmt"
	"io"
	"time"
	"crypto/md5"
)

type User struct {
	UserName   		string    `xorm:"not null  pk varchar(255)" json:"name"`
	UserPassword	string    `xorm:"not null  VARCHAR(255)" json:"password"`
	LoginTime  		time.Time `xorm:"DATETIME" json:"login_time"`
	//Role 			int 	  `xorm:"not null  NUMBER(255)" json:"role"`

}

func	UserCheck(name string)int {
	result,err:=Db.Exist(&User{
		UserName: name,
	})
	if err!=nil{
		panic(err)
	}

	if result {
		return utils.SUCCESS
	}
	return utils.ERROR
}
func (u *User)Encode()  {
	u.UserPassword=EnPassWord(u.UserPassword)
}
func 	EnPassWord(password string)string  {
	w:=md5.New()
	io.WriteString(w,password)
	bw:=w.Sum(nil)
	newpwd:=hex.EncodeToString(bw)
	return newpwd
}
func	InsertUser(data *User)int {
	data.Encode()
	result,err:=Db.InsertOne(data)
	if err!=nil{
		panic(err)
	}
	if result>0 {
		return utils.SUCCESS
	}
	return utils.ERROR
}

func	DeleteUser(name string)int{
	var user User
	result,err:=Db.Where("user_name = ?",name).Delete(&user)
	if err!=nil{
		panic(err)
	}

	if result > 0{
		return utils.SUCCESS
	}
	return utils.ERROR
}
func LoginCheck(username string ,password string)int{
	var user User
	fmt.Println("123"+username)
	code := UserCheck(username)
	if code == utils.ERROR{
		return utils.USER_NOT_EXIST
	}
	Db.Where("user_name = ?",username).Get(&user)
	if user.UserPassword!=EnPassWord(password){
		return utils.PASSWORD_ERROR
	}
	return utils.SUCCESS
}