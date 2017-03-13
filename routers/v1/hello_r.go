package v1





import (
		
	hello "simple-api/controllers/v1"
	"gopkg.in/gin-gonic/gin.v1"	

)


func SetHelloRoutes(router *gin.RouterGroup) {
	
	router.GET("/hello" , hello.HelloUser)
}
