package appid

import (
	"gopkg.in/gin-gonic/gin.v1"
	// token "simple-api/services/middleware/token"
	// appid "simple-api/services/middleware/appid"
		"time"
	"simple-api/core/models"
	// "fmt"
	 // "encoding/json"
	 // "errors"
)



func returnErrorNotValid(c *gin.Context) {
	
	var content gin.H;
	var listError [] models.ErrorModel = nil
	var statuscode 	int
	var endpoint	string = c.Request.URL.String()
	var method		string =c.Request.Method
	t :=time.Now();
	 
		statuscode =401
			 
			var tempEr models.ErrorModel
			tempEr.ErrorCode 	= 4014	
			tempEr.Text 		= "Invalid Token"
			tempEr.Hints 		= "Make Sure include the valid X-Auth-Token Header"
			tempEr.Info			= "visit support.example.com"
			listError 			= append(listError,tempEr)
			 

		content = models.NewResponse(t,statuscode,"Fail",nil,listError,endpoint,method);	
		
		c.JSON(statuscode,content)
		c.Abort()
	
}


func ValidateAppIDMiddleWare() gin.HandlerFunc {
	return func (c *gin.Context) {

		/*
		*
		*Currently Empty, For more fucntions to validate the X-App-ID header
		*
		**/


			c.Next()		
		
 		
 	}
}
