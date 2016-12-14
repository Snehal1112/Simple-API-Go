package users


import (
		
	"simple-api/controllers/users"
	"github.com/gin-gonic/gin"	

)


func SetGetAllTaskByProject(router *gin.RouterGroup) {
	
	router.GET("/kasisemua/:project_id", users.GetAllTaskByProject)
}
