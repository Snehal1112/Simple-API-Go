package users

import (

	"github.com/gin-gonic/gin"	
)


func GetAllTaskByProject( c *gin.Context) {
	
content := gin.H{
            "result": "Success",
            "title": "USER",
            "content": "Yeayyyy",
        }
        c.JSON(200, content)

}
