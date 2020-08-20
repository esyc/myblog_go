package utils

import (
	"os"
	"encoding/json"
	"bufio"
)

type Config struct {
	BlogName string	`json:"blog_name"`
	BlogMode string	`json:"blog_mode"`
	BlogHost string	`json:"blog_host"`
	BlogPort string	`json:"blog_port"`
	BlogJwt string 	`json:"jwt"`
	Database DatabaseConfig `json:"database"`
}
type DatabaseConfig struct {
	Driver 			string `json:"driver"`
	User 			string `json:"user"`
	Password		string `json:"password"`
	Host			string `json:"host"`
	Port			string `json:"port"`
	DatabaseName 	string `json:"database_name"`
	Charset      	string `json:"charset"`
	Show			bool 	`json:"show"`
}
var cfg *Config=nil

func GetConfig(path string)(*Config,error)  {
	file,err :=os.Open(path)
	if err!=nil{
		panic(err)
	}
	defer file.Close()

	reader:=bufio.NewReader(file)
	decoder:=json.NewDecoder(reader)
	if err=decoder.Decode(&cfg);err !=nil{
		return nil,err
	}
	return cfg,nil
}