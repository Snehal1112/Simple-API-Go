package appid





import (
	"gopkg.in/gin-gonic/gin.v1"
	// services "simple-api/services/middleware/token"
		"time"
	"simple-api/core/models"
	// "fmt"
)


func AppIDMiddleWare() gin.HandlerFunc {

	return func (c *gin.Context) {

		appID := c.Request.Header.Get("X-App-ID")
		
		if 	appID == "oktaskmobile"  ||		
			appID == "oktaskweb" ||		
			appID == "postman" {
			c.Next()				
		}else{
			var content gin.H
			var listError [] models.ErrorModel = nil
			var statuscode 	int
			var endpoint	string = c.Request.URL.String()
			var method		string =c.Request.Method
			t :=time.Now()
 	
	
	

	 
			statuscode =403
			 
			var tempEr models.ErrorModel
			tempEr.ErrorCode 	= 4031	
			tempEr.Text 		= "Invalid Device"
			tempEr.Hints 		= "Make Sure you are calling from the registered device"
			tempEr.Info			= "visit support.example.com"
			listError 			= append(listError,tempEr)
			 

			content = models.NewResponse(t,statuscode,"Fail",nil,listError,endpoint,method);	
			
			c.JSON(statuscode,content)
			c.Abort()
	
		}	
 
 	}
}
 