package users


import (
		
	"simple-api/controllers/users"
	"github.com/gin-gonic/gin"	

)


func SetGetAllTaskByProject(router *gin.RouterGroup) {
	

	/**
 * @api {get} /task/kasisemua/:project_id Get All Task By Project Id
 * @apiGroup Users
 * @apiHeader {application/json} Content-Type Accept application/json
 * @apiParam {Integer} Project_Id Project ID
 * @apiParamExample Input
 * http://localhost:9090/task/kasisemua/1
 * @apiSuccess {Object} content Data
 * @apiSuccess {String} content.title_task Judul Task
 * @apiSuccess {String} content.difficulty Tingkat Kesukaran Task
 * @apiSuccess {String} content.urgency Tingkat Kepentingan Task
 * @apiSuccess {String} content.desc_task Deskripsi Task
 * @apiSuccess {String} content.pemberi Pemeberi Task
 * @apiSuccess {String} content.penerima Penerima Task
 * @apiSuccessExample {json} Success
	{
	  "content": [
	    {
	      "title_task": "Membuat Layout Login",
	      "difficulty": "4",
	      "urgency": "5",
	      "desc_task": "Create Login Layout",
	      "pemberi": "Iman",
	      "penerima": "Eminarty "
	    }
	  ],
	  "status code": 200,
	  "success": true
	}
 * @apiErrorExample {json} List error
 *    HTTP/1.1 500 Internal Server Error
 */
	router.GET("/kasisemua/:project_id", users.GetAllTaskByProject)
}
