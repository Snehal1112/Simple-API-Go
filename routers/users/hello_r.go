package users


import (
		
	"simple-api/controllers/users"
	"github.com/gin-gonic/gin"	

)


func SetHelloRoutes(router *gin.RouterGroup) {
	
	router.GET("/hello" , users.HelloUser)
}
