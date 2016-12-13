package users


import (
	"github.com/gin-gonic/gin"	
	"net/http"
	"simple-api/core/models"
	. "simple-api/settings"
)


func InitRoutes(g *gin.RouterGroup)  {
	g.Use(RoleValidationMiddleWare());
	SetHelloRoutes(g)
}

 func RoleValidationMiddleWare() gin.HandlerFunc{
 	return func (c *gin.Context) {
		decodedToken := c.MustGet("decodedToken").(*models.ClaimsData)
		token := c.Request.Header.Get("X-Auth-Token")

		akun,err := models.FindOkAkun(DBCon,decodedToken.AkunId)


			if decodedToken.Role != "user"  {
		 			content:= gin.H{
							"success" : false,
							"message" : "Role Tidak sesuai",
							"statuscode": http.StatusForbidden, 
							}
				c.JSON(http.StatusForbidden, content)
				 c.Abort()
			}



			if  akun.Token != token {
		 			content:= gin.H{
							"success" : false,
							"message" : "Token Tidak Sama",
							"statuscode": http.StatusForbidden, 
							}
				c.JSON(http.StatusForbidden, content)
				 c.Abort()
			}


			if err!= nil {
		 			content:= gin.H{
							"success" : false,
							"message" : err,
							"statuscode": http.StatusForbidden, 
							}
				c.JSON(http.StatusForbidden, content)
				 c.Abort()
			}

 		c.Next();
 	}
 }
