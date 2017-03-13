package main

import (
	"net/http"
	// "github.com/fvbock/endless"
	"fmt"
	"time"
	"database/sql"
	// _ "github.com/lib/mq"
	// _ "github.com/go-sql-driver/mysql"
	// services "simple-api/services/token"
	// "gopkg.in/nullbio/null.v6"
	// "github.com/vattle/sqlboiler/boil"
	"gopkg.in/inconshreveable/log15.v2"
	mysql "simple-api/services/mysql"
	mongodb "simple-api/services/mongodb"
	// elastic "simple-api/services/elastic"
	//	"simple-api/routers/drivers"
	//	"simple-api/routers/users"
	//	"simple-api/routers/driver_auth"
	//	"simple-api/routers/user_auth"
	//	"simple-api/routers/apidoc"
	"simple-api/routers/v1"
	"gopkg.in/gin-gonic/gin.v1"
	"github.com/vattle/sqlboiler/boil"
	"simple-api/core/models"
	cors "simple-api/middlewares/cors"
	header "simple-api/middlewares/header"

	"gopkg.in/mgo.v2"
)

var log = log15.New();

var router  *gin.Engine



var (
	DBCon *sql.DB
	Session *mgo.Session
)

func init () {
	
	DBCon= mysql.CheckDb()
	// elastic.CheckElastic()
	router = gin.New()
	//mongodb.CheckMongoDB()
	
	router.Use(header.HeaderMiddleWare())
	router.Use(cors.CORSMiddleware())
	router.NoRoute(noRouteHandler())
	

	v1Api := router.Group("/v1")
	v1.InitRoutes(v1Api)



}


func main () {
boil.DebugMode = true
	
	//defer Session.Close();
	
	
	log.Info("Server Running on ", "port" , 9090)		
	http.ListenAndServe(":9090", router)
	// endless.ListenAndServe(":9090", router)

}



	
func noRouteHandler() gin.HandlerFunc{


	return  func(c *gin.Context) {


	var content gin.H;
	var listError [] models.ErrorModel = nil
	var statuscode 	int
	var endpoint	string = c.Request.URL.String()
	var method		string = c.Request.Method
	t :=time.Now();

	statuscode =404
	var tempEr models.ErrorModel
	tempEr.ErrorCode 	= 4041	
	tempEr.Text 		= "Not Found"
	tempEr.Hints 		= "Routes In Valid. You enter on invalid Page/Endpoint"
	tempEr.Info			= "visit http://api.oktravel.com/docs/akun#172401"
	listError 			= append(listError,tempEr)
	
	content = models.NewResponse(t,statuscode,"Fail",nil,listError,endpoint,method);	
	c.JSON(statuscode,content)

	}
    

}


func SyncDataMongo() {
	

	mongodb.DeleteTaskRefCollections()
	
	TaskData,err:= models.Oktasks(DBCon).All()

	fmt.Println(err)
	for _,v:= range TaskData{
		mongodb.InsertDataTask(*v)	
	}
	
	
		
}
