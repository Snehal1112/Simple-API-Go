package services

import (
	"gopkg.in/gin-gonic/gin.v1"
	"github.com/dgrijalva/jwt-go"
	"gopkg.in/inconshreveable/log15.v2"
	"simple-api/core/models"
	"io/ioutil"
	"errors"
	"fmt"
	// "simple-api/services/redis"
	// elastic "simple-api/services/elastic"
	"time"
	// . "simple-api/services/mysql"
	// "fmt"
	// "strconv"
	// "bytes"
	// "simple-api/libs"
	// "errors"
	// "strings"
	// "github.com/vattle/sqlboiler/queries/qm"

)

var log = log15.New()

func ValidateToken(c *gin.Context) {
	token := c.Request.Header.Get("X-Auth-Token")

	var (
		status bool
		msg interface{}
		) 

	if token !=""{
		status,msg = checkToken(token) ;	
	}
	

	var content gin.H;
	var listError [] models.ErrorModel = nil
	var statuscode 	int
	var endpoint	string = c.Request.URL.String()
	var method		string =c.Request.Method
	t :=time.Now();
 	
	
	

	if  !status ||  token  == ""  {

		statuscode =401
			
		// if token == "" || !status{
			
			var tempEr models.ErrorModel
			tempEr.ErrorCode 	= 4014	
			tempEr.Text 		= "Invalid Token"
			tempEr.Hints 		= "Make Sure include the valid X-Auth-Token Header"
			tempEr.Info			= "visit support.example.com"
			listError 			= append(listError,tempEr)
			
		// }



		content = models.NewResponse(t,statuscode,"Fail",nil,listError,endpoint,method);	
		
		c.JSON(statuscode,content)
		c.Abort()
		// c.AbortWithStatus(401)
	}else{

		c.Set("decodedToken",msg);
		c.Next()
	}
 

}





func checkToken(tokenInput string ) (bool, interface{} ) {
	//TODO : Verify the token signature with our secret key, and Decode the payload to give to next routers 
	
	token  ,err := jwt.ParseWithClaims(tokenInput,  &models.ClaimsData{}, func(token *jwt.Token) (interface {} , error) {
				key, err1 := ioutil.ReadFile("services/ninja.naruto")

				
    			if err1 != nil {
        			return nil, errors.New("private key could not be loaded")
    			}
    			    			
    			return key, nil
		})


	if err !=nil{
		log.Error("Error Verify", "err", err);
	}

	// claims ,ok :=token.Claims.(*models.ClaimsData)
	// if !ok {
	// 	log.Error("ERROR" , "err" , ok);
	// }

	// fmt.Println("HOALS")
	fmt.Println(token.Valid)

	return token.Valid, token.Claims;
}

 


func GetNewToken(data interface{}) (bool,string) {

	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["data"] = data
	// claims["exp"] = time.Now().Add(time.Hour *72).Unix()
	token.Claims = claims
	key, err := ioutil.ReadFile("services/ninja.naruto")

	if err != nil {
	    log.Error("Error","private key could not be loaded",err)
	    return false,"failed";
    }else{
    	
    	tokenString , err := token.SignedString(key)    	

    	if err == nil {
    		return true,tokenString;
    	}else{
    		return false,"failed"
    	}
    		
    }

 }
	
