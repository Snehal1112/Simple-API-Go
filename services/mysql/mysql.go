package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"gopkg.in/inconshreveable/log15.v2"
)



var (
	DBCon *sql.DB
) 

var log = log15.New();

// var db *sql.DB
var err error
func init() {
 ConnectDB();	
}

func ConnectDB() {
	DBCon,err =sql.Open("mysql","root:15pstr@tcp(127.0.0.1:3306)/ok_data?parseTime=true")
	// DBCon,err =sql.Open("mysql","user:user@tcp(127.0.0.1:3306)/OkData?parseTime=true")
}


func CheckDb() *sql.DB  {
	
	if   pinger  := DBCon.Ping() ; err != nil || pinger!=nil  {
  			log.Error("FAILDE TO CONNECT", "err" , err)
  			//TODO : make a retry connect
		}
	
	return DBCon
}