package v1


import (
	"gopkg.in/gin-gonic/gin.v1"
	// token "simple-api/middlewares/token"
	
	appid "simple-api/middlewares/appid"
	
		// "time"
	// "simple-api/core/models"
	// "fmt"
)




func InitRoutes(g *gin.RouterGroup)  {
	
	g.Use(appid.AppIDMiddleWare())
	SetHelloRoutes(g)
	SetAuthRoutes(g)
	/*
	* UnComment when all already ready ( finished)
	*
	*g.Use(token.TokenAuthMiddleWare()) 
	g.Use(appid.ValidateAppIDMiddleWare())*/
	SetTaskRoutes(g)
	SetUserRoutes(g)
}


