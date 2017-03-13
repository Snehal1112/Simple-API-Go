package v1





import (
		
	// "simple-api/controllers/auth/driver"
	"gopkg.in/gin-gonic/gin.v1"	

)


func SetApiDocRoutes(router *gin.RouterGroup) {
	router.Static("/docs", "public/apidoc")
	
}