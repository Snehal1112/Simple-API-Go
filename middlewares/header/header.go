package header


import (
	
	"gopkg.in/gin-gonic/gin.v1"
	// "fmt"
	"time"
	"simple-api/core/models"
)


 func HeaderMiddleWare() gin.HandlerFunc{
 	return func (c *gin.Context) {
 		


 		heder:= c.Request.Header

 		contentType:= heder.Get("Content-Type")

 		if contentType != "application/json"{

 			
			var content gin.H;
			var listError [] models.ErrorModel = nil
			var statuscode 	int
			var endpoint	string = c.Request.URL.String()
			var method		string =c.Request.Method
			t :=time.Now();
			 
				statuscode =415
					 
					var tempEr models.ErrorModel
					tempEr.ErrorCode 	= 4151	
					tempEr.Text 		= "Unsupported Media Type"
					tempEr.Hints 		= "Only Accept JSON Data type"
					tempEr.Info			= "visit support.example.com"
					listError 			= append(listError,tempEr)
					 

				content = models.NewResponse(t,statuscode,"Fail",nil,listError,endpoint,method);	
				
				c.JSON(statuscode,content)
				c.Abort()
	 	}


	 		c.Next()

 	}

 }


 func errorMismatchDataType() {
 	
 }