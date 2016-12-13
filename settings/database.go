package settings

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

	
	DBCon,err =sql.Open("mysql","root:15pstr@tcp(127.0.0.1:3306)/OkPortofolio?parseTime=true")
	// DBCon,err =sql.Open("mysql","user:user@tcp(127.0.0.1:3306)/OkData?parseTime=true")
	

}


func CheckDb() {
	
	if err != nil {
  			log.Error("FAILDE TO CONNECT", "err" , err)
		}
	if pinger  := DBCon.Ping() ; pinger!=nil {
			log.Error("FAILDE TO CONNECT", "err" , pinger)
	}	

	// return db
}