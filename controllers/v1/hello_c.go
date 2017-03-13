package v1




import (
	"gopkg.in/gin-gonic/gin.v1"
)



func HelloUser(c *gin.Context) {
			content := gin.H{
            "result": "Success",
            "title": "USER",
            "content": "Success To Hello",
        }
        c.JSON(200, content)

}

