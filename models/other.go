package models

import (
	"crcblog/utils"
)

type Zan struct {
	Number int `xorm:" int" json:"zan"`
}
type People struct {
	Number int `xorm:" int" json:"people"`
}
type Message struct {
	Msg string`xorm:"varchar(255)" json:"msg"`
}
func GetPeo() (People,int) {
	var p People
	var ps []People
	data,_:= Db.Count(&p)
	if data == 0{
		p.Number = 0
	} else{
		_ =Db.Table("People").Find(&ps)
		p.Number = ps[0].Number
	}
	return p,utils.SUCCESS
}
func UpdatePeo()  {
	var p People
	var ps []People
	data,_:= Db.Count(&p)
	if data == 0{
		p.Number = 1
		Db.InsertOne(&p)
	} else{
		_ =Db.Table("People").Find(&ps)
		p.Number = ps[0].Number + 1
		Db.Update(&p)
	}
}
func GetZan() (Zan,int) {
	var z Zan
	var zs []Zan
	data,_:= Db.Count(&z)
	if data == 0{
		z.Number = 0
	} else{
		_ =Db.Table("Zan").Find(&zs)
		z.Number = zs[0].Number
	}
	return z,utils.SUCCESS
}
func UpdateZan()  {
	var z Zan
	var zs []Zan
	data,_:= Db.Count(&z)
	if data == 0{
		z.Number = 1
		Db.InsertOne(&z)
	} else{
		_ =Db.Table("Zan").Find(&zs)
		z.Number = zs[0].Number + 1
		Db.Update(&z)
	}
}
func GetMessage() ([]Message,int) {
	var m []Message
	err :=Db.Table("Message").Find(&m)
	code := utils.SUCCESS
	if err !=nil{
		panic(err)
		code = utils.ERROR
	}
	return m,code
}
func InsertMessage(msg Message)  {
	var m Message
	m.Msg = msg.Msg
	Db.InsertOne(&m)
}