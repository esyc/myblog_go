package utils

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)
type OrmEngine struct {
	*xorm.Engine
}
//初始化数据库xorm+mysql
func InitOrmDb(cfg *Config)(*OrmEngine,error){
	database :=cfg.Database
	connection:=database.User+":"+database.Password+"@tcp("+database.Host+":"+ database.Port+")/"+database.DatabaseName+"?charset="+database.Charset
	engine,err:=xorm.NewEngine(database.Driver,connection)
	if err!=nil{
		return nil,err
	}
	engine.ShowSQL(database.Show)
	orm:=new(OrmEngine)
	orm.Engine=engine
	return orm,nil
}