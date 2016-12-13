package main

import (
	"net/http"
	// "fmt"
	// "time"
	// "database/sql"
	// _ "github.com/lib/mq"
	// _ "github.com/go-sql-driver/mysql"
	"simple-api/services"
	// "gopkg.in/nullbio/null.v6"
	// "simple-api/services/models"
	// "github.com/vattle/sqlboiler/boil"
	"gopkg.in/inconshreveable/log15.v2"
	. "simple-api/settings"
	"simple-api/routers/users"
	"simple-api/routers/user_auth"
	"simple-api/routers/apidoc"
	"github.com/gin-gonic/gin"
)

var log = log15.New();

var router  *gin.Engine


func init() {

	CheckDb()
	router = gin.New()
	
	// router.StaticFile("/ninjaApi", "./public/apidoc/index.html")
	router.Use(CORSMiddleware())
	authUserApi := router.Group("/authuser")
	user_auth.InitRoutes(authUserApi)
	docAPI := router.Group("/ninja");
	apidoc.InitRoutes(docAPI)
	router.Use(TokenAuthMiddleWare())
	userApi := router.Group("/task")
	users.InitRoutes(userApi)

}


func main () {

	log.Info("Server Running on ", "port" , 9090)		
	http.ListenAndServe(":9090", router)

}


func CORSMiddleware() gin.HandlerFunc {

     return func(c *gin.Context) {

         c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
         c.Writer.Header().Set("Access-Control-Max-Age", "86400")
         c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, PUT, DELETE, UPDATE")
         c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
         c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
         c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

         if c.Request.Method == "OPTIONS" {
             c.AbortWithStatus(204)
         } else {
             c.Next()
         }
     }
 }

 func TokenAuthMiddleWare() gin.HandlerFunc{
 	return func (c *gin.Context) {
 		services.ValidateToken(c);
 		c.Next();
 	}

 }

