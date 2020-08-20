package models

import (
	"crcblog/utils"
	"github.com/go-xorm/xorm"
)
var Db *xorm.Engine
func InitDatabase(cfg *utils.Config) {
	result,err:=utils.InitOrmDb(cfg)

	if err!=nil{
		panic(err)
	}
	err = result.Sync2(new(User),new(Category),new(Post),new(Zan),new(Message),new(People))
	if err!=nil{
		panic(err)
	}
	Db=result.Engine
}
