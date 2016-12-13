package user_auth



import (
		
	"simple-api/controllers/auth/user"
	"github.com/gin-gonic/gin"	

)


func SetAuthRoutes(router *gin.RouterGroup) {
	


/**
 * @api {post} /authuser/login Login
 * @apiGroup Users
 * @apiHeader {application/json} Content-Type Accept application/json
 * @apiParam {String} email User email
 * @apiParam {String} password User Password
 * @apiParamExample {json} Input
 *    {
 *      "email": "your email",
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

	router.POST("/login" , user.Login)
	// router.POST("/register" , users.Login)
	// router.POST("/logout" , users.Login)
	// router.POST("/refreshtoken" , users.Login)
	// router.POST("/login" , users.Login)

	
}

