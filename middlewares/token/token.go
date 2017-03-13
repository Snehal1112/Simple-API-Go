package token


import (
	
	"gopkg.in/gin-gonic/gin.v1"
	services "simple-api/services/token"
)


 func TokenAuthMiddleWare() gin.HandlerFunc{
 	return func (c *gin.Context) {
 		services.ValidateToken(c);
 		c.Next();
 	}

 }