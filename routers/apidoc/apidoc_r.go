package apidoc


import (
		
	// "okbackend/controllers/auth/driver"
	"github.com/gin-gonic/gin"	

)


func SetApiDocRoutes(router *gin.RouterGroup) {
	router.Static("/docs", "public/apidoc")
	
}