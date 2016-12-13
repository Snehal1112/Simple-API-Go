package user

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"simple-api/core/models"
	. "simple-api/settings"
	"simple-api/services"
	"gopkg.in/inconshreveable/log15.v2"
	. "github.com/vattle/sqlboiler/queries/qm"

	// "database/sql"
	// _ "github.com/lib/mq"
	// _ "github.com/go-sql-driver/mysql"
	// "gopkg.in/nullbio/null.v6"
)



type LoginModel struct {
	username  string ` form:"username" json : "username" binding: "required"`
	Password  string ` form:"password" json : "password" binding: "required"`
}



var log = log15.New();

func Login(c *gin.Context) {

	
	// var Json LoginModel 
	var Form LoginModel
	c.Bind(&Form)


	if Form.Password== "" {

		content:= gin.H{
			"success" : false,
			"message" : "No Password Input",
			
		}
		c.JSON(http.StatusBadRequest, content)
		return
	}
	if Form.username== "" {

		content:= gin.H{
			"success" : false,
			"message" : "No username Input",
			
		}
		c.JSON(http.StatusBadRequest, content)
		return
	}

	
	akun, err := models.OkAkuns(DBCon, Where("username = ?", Form.username), And("role=?", "user")).One()

	if err !=nil {
			content:= gin.H{
				"success" : false,
				"message" : "Not Found",
				"errorcode": http.StatusNotFound,
			}
			log.Error("error", "err", err);
		c.JSON(http.StatusNotFound, content);
		return

	}else {
			if akun.Password != Form.Password	 {
				content:= gin.H{
					"success" : false,
					"message" : "Wrong Password", 
					"errorcode": http.StatusNotFound,
				}
				c.JSON(http.StatusNotFound, content)
				return
			}else {

				user,err := akun.AkunOkProfiles(DBCon).One()

				if err != nil{
					log.Error("Error", "err", err)

					content:= gin.H{
							"success" : false,
							"message" : "Not Found", 
					}
					c.JSON(http.StatusNotFound, content)
					return		
				}
				
					data := &models.JwtData{

						Name 		: user.FirstName + " " + user.LastName,
						Role		: akun.Role,
						UserName 	: akun.Username,
						AkunId 		: akun.IDAkun,
						ProfileId 	: user.IDProfile,
					}

				 if	status,token := services.GetNewToken(data) ; status ==true {
			 		var authenticated  models.LoginResponse;
				 	go insertToDb(token,akun)
				 	authenticated = models.LoginResponse{Token:token, Success : true , Message: "Login Successfully" , Statuscode : http.StatusOK}

				 	   content:= gin.H{
							"authenticate" :authenticated,
						}
						c.JSON(http.StatusOK, content)
						return
				 }else{
					 	content:= gin.H{
						// "token" :token,
						"success" : false,
						"message" : "Internal Server Error",
						"statuscode": http.StatusInternalServerError, 
					}
					c.JSON(http.StatusInternalServerError, content)
					return	
				 }
			}	
		 

	}

}

func insertToDb(token string, akun *models.OkAkun) {
	
	akun.Token = token;

	err:= akun.Update(DBCon)
	if err !=nil{
		fmt.Println(err)
	}
}



