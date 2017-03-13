package models


import (
 	"gopkg.in/gin-gonic/gin.v1" 
    "time"
    // "fmt"
)


const (
	 DATE_FORMAT = "2006-01-02 15:04:05"
  	 DATE_REGEX = "((19|20)\\d\\d)-(0?[1-9]|1[012])-(0?[1-9]|[12][0-9]|3[01])"	
  	 // match, err := regexp.MatchString("pattern", "string")
)



type ResponseModel struct{

	t 				time.Time
	statuscode 		int
	message 		string
  	data 			interface{}
  	listError 		[]ErrorModel
  	endpoint 		string
  	method 			string
}



func NewResponse( 
	t 			time.Time,
	statuscode 		int,
	message 		string,
  	data 			interface{},
  	listError 		[]ErrorModel,
  	endpoint 		string,
  	method 			string)gin.H {	


	content:= gin.H{
			"timestamp"	: t.Format(DATE_FORMAT),	 
			"statuscode": statuscode,
			"message" 	: message,
			"data"		: data,
			"error"		: listError,
			"endpoint"	: endpoint,
			"method"	: method,
		}

return content;
}