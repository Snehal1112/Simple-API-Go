package user_auth



import (
"github.com/gin-gonic/gin"	
)


func InitRoutes(g *gin.RouterGroup)  {
	SetAuthRoutes(g)
}

