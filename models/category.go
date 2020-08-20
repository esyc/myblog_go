package models

import (
	"crcblog/utils"
)

type Category struct{
	Id int			`xorm:"not null autoincr pk int " json:"id"`
	CateName string `xorm:"not null varchar(255)" json:"name"`
}
func	CateCheck(name string)int {
	result,err:=Db.Exist(&Category{
		CateName: name,
	})
	if err!=nil{
		panic(err)
	}

	if result {
		return utils.SUCCESS
	}
	return utils.ERROR
}
func GetAllCate()[]Category {
	var cates []Category
	err:=Db.Find(&cates)
	if err!=nil{
		panic(err)
	}
	return cates
}
func GetCatefromId(id int)Category {
	cate := new(Category)
	_,err:=Db.Where("id = ?",id).Get(cate)
	if err!=nil{
		panic(err)
	}
	return *cate
}
func InsertCate(data *Category)int{


	result,err:=Db.InsertOne(data)
	if err!=nil{
		panic(err)
	}
	if result>0 {
		return utils.SUCCESS
	}
	return utils.ERROR
}
func DeleteCate(id int)int{
	var cate Category
	result,err:=Db.Where("id = ?",id).Delete(&cate)
	if err!=nil{
		panic(err)
	}

	if result > 0{
		return utils.SUCCESS
	}
	return utils.ERROR
}
func EditCate(id int ,data *Category)int{
	var cate Category
	cate.Id=data.Id
	cate.CateName=data.CateName
	_,err:=Db.Where("id = ?",id).Update(&cate)
	if err!=nil{
		return utils.ERROR
	}
	return utils.SUCCESS
}