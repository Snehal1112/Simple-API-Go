package users

import (

	"gopkg.in/gin-gonic/gin.v1"	
	// "simple-api/core/models"
	mysql "simple-api/services/mysql"
	"github.com/vattle/sqlboiler/queries"
	"fmt"
)



type ResultSet struct {

 // NamaProject string    	`boil:"nama_project" json:"nama_project" toml:"nama_project" yaml:"nama_project`
 // Desc  		 string		`boil:"desc" json:"desc" toml:"desc" yaml:"desc`
 // GitURL		 string		`boil:"git_url" json:"git_url" toml:"git_url" yaml:"git_url`
 TitleTask	 string		`boil:"title_task" json:"title_task" toml:"title_task" yaml:"title_task`
 Difficulty	 string		`boil:"difficulty" json:"difficulty" toml:"difficulty" yaml:"difficulty`
 Urgency	 string		`boil:"urgency" json:"urgency" toml:"urgency" yaml:"urgency`
 DescTask	 string		`boil:"desc_task" json:"desc_task" toml:"desc_task" yaml:"desc_task`	
 Pemberi	 string		`boil:"pemberi" json:"pemberi" toml:"pemberi" yaml:"pemberi`
 Penerima	 string		`boil:"penerima" json:"penerima" toml:"penerima" yaml:"penerima`

}


func GetAllTaskByUser( c *gin.Context) {
	
 idProject := c.Param("project_id");

var resultset [] ResultSet
// I Use Call Procedure From Database; Look on Mysql For detailed the Query
err := queries.Raw(mysql.DBCon,`CALL get_all_task_by_project(?)`, idProject).Bind(&resultset)
			
if err !=nil {

		fmt.Println(err);
		content := gin.H{
            "status_code": 500,
            "message": "Internal Server Error",
        }

        c.JSON(500, content)

}else {
	content := gin.H{
            "success": true,
            "status_code": 200,
            "content": resultset,
        }

        c.JSON(200, content)

}			


}
