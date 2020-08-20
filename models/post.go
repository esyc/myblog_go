package models

import (
	"crcblog/utils"
	"time"
)

type Post struct {
	Cate Category
	PostId int			`xorm:"not null  autoincr int pk" json:"postid"`
	Title string		`xorm:"not null  varchar(255)" json:"title"`
	CId int				`xorm:"int" json:"cateid"`
	Description string	`xorm:"longtext" json:"description"`
	Content string 		`xorm:"not null  longtext" json:"content"`
	//Img string			`xorm:"varchar(255)" json:"img"`
	PostTime time.Time  `xorm:"DateTime" json:"time"`
}

func InsertPost(data *Post)int  {
	result,err:=Db.InsertOne(data)
	if err!=nil{
		panic(err)
	}
	if result>0 {
		return utils.SUCCESS
	}
	return utils.ERROR
}
func GetPostByNum(pageSize int,pageNum int)[]Post {
	var posts =make([]Post,0)
	err :=Db.Limit(pageSize,pageSize*(pageNum-1)).Find(&posts)
	for i := 0; i < len(posts);i++{
		posts[i].Cate = GetCatefromId(posts[i].CId)
	}

	if err!=nil{
		panic(err)
	}
	return posts
}
func GetAllPost()[]Post  {
	var posts =make([]Post,0)
	err :=Db.Table("post").Join("INNER", "category", "post.c_id = category.id").
		Find(&posts)

	for i := 0; i < len(posts);i++{
		posts[i].Cate =GetCatefromId(posts[i].CId)
	}
	if err!=nil{
		panic(err)
	}
	return posts
}
func GetPostFromCate(id int)[]Post {
	var posts []Post
	err :=Db.Where("c_id = ?",id).Find(&posts)
	for i := 0; i < len(posts);i++{
		posts[i].Cate =GetCatefromId(posts[i].CId)
	}
	if err!=nil {
		panic(err)
	}
	return posts
}
func GetOnePost(id int)(Post,int)  {
	var post Post
	result,err:= Db.Where("post_id = ?",id).Get(&post)
	if err !=nil{
		panic(err)
	}
	code := utils.SUCCESS
	if !result {
		code = utils.POST_NOT_FOUND
	}
	post.Cate=GetCatefromId(post.CId)
	return post,code
}

func EditPost(id int,data *Post)int  {
	var post Post
	post.Title=data.Title
	post.CId=data.CId
	post.Description=data.Description
	post.Content=data.Content
	_,err:=Db.Where("post_id = ?",id).Update(&post)
	if err!=nil{
		return utils.ERROR
	}
	return utils.SUCCESS
}

func DeletePost(id int)int{
	var post Post
	result,err:=Db.Where("post_id = ?",id).Delete(&post)
	if err!=nil{
		panic(err)
	}

	if result > 0{
		return utils.SUCCESS
	}
	return utils.ERROR
}
func GetPostNum()int64  {
	var p Post
	data,_:= Db.Count(&p)
	return data
}