package users

import (
	"github.com/gin-gonic/gin"
)



func HelloUser(c *gin.Context) {
			content := gin.H{
            "result": "Success",
            "title": "USER",
            "content": "Success To Hello",
        }
        c.JSON(200, content)

}

