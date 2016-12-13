package services

import (
	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
	"gopkg.in/inconshreveable/log15.v2"
	// "time"
	"simple-api/core/models"
	"io/ioutil"
	"errors"
	// "fmt"
)

var log = log15.New()

func ValidateToken(c *gin.Context) {
	token := c.Request.Header.Get("X-Auth-Token")

	if token  == ""{
		content := gin.H{
			"error": 401,
			"message" : "No Token Provided",
			"success" : false,
		}

		c.JSON(401,content)
		c.AbortWithStatus(401)
	}else if  status,msg := checkToken(token) ; status ==true{
		// fmt.Println(msg.Role);	
		c.Set("decodedToken",msg);
		c.Next()
	}else {
		content := gin.H{
			"error": 401,
			"message": "No Token Provided",
			"success" : false,
		}
		c.JSON(401,content)
		c.AbortWithStatus(401)
	}

}





func checkToken(tokenInput string ) (bool, *models.ClaimsData ) {
	//TODO : Verify the token signature with our secret key, and Decode the payload to give to next routers 
	
	token  ,err := jwt.ParseWithClaims(tokenInput,  &models.ClaimsData{}, func(token *jwt.Token) (interface {} , error) {
				key, err := ioutil.ReadFile("services/ninja.naruto")

				
    			if err != nil {
        			return nil, errors.New("private key could not be loaded")
    			}
    			    			
    			return key, nil
		})


	if err !=nil{
		log.Error("Error Verify", "err", err);
	}

	claims ,ok := token.Claims.(*models.ClaimsData)
	if !ok {
		log.Error("ERROR" , "err" , ok);
	}

	return token.Valid, claims;
}




func GetNewToken(data *models.JwtData) (bool,string) {

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
	
