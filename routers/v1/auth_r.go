package v1






import (
		
	auth "simple-api/controllers/v1/auth"
	"gopkg.in/gin-gonic/gin.v1"	

)


func SetAuthRoutes(router *gin.RouterGroup) {
	


/**
 * @api {post} /v1/auth/login Login
 * @apiGroup Users
 * @apiHeader {application/json} Content-Type Accept application/json
 * @apiParam {String} username User username
 * @apiParam {String} password User Password
 * @apiParamExample {json} Input
 *    {
 *      "username": "your username",
 *		"password"	 : "your password"		
 *    }
 * @apiSuccess {Object} authenticate Response
 * @apiSuccess {Boolean} authenticate.success Status
 * @apiSuccess {Integer} authenticate.statuscode Status Code
 * @apiSuccess {String} authenticate.message Authenticate Message
 * @apiSuccess {String} authenticate.token Your JSON Token
 * @apiSuccessExample {json} Success
 *    {
 *		"authenticate": { 	
 *	 	 	 "statuscode": 200,
 *     		 "success": true,
 *      	 "message": "Login Successfully",
 *	 	     "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ"
 *    		}
 *	  }
 * @apiErrorExample {json} List error
 *    HTTP/1.1 500 Internal Server Error
 */

	router.POST("/auth/login" , auth.Login)
	

	
}

