package controllers

import (
	"net/http"

	"example.com/jwt/initializers"
	"example.com/jwt/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SignUp( c *gin.Context ) {
	//이메일 패스워드 req.body
	var body struct{
		Email string
		Password string
	}

	if c.Bind(&body) !=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"Faild to read Bdody",
		})

		return
	}

	//hash password
	hash,err := bcrypt.GenerateFromPassword([]byte(body.Password),10)
	
	if err!= nil {
		c.JSON(http.StatusBadRequest, gin.H{
            "error":"Faild to hash password",
        })

        return
	}
	
	//create user
	user := models.User{Email: body.Email, Password: string(hash)}

	result := initializers.DB.Create(&user)

	if result.Error!= nil {
		c.Status(400)
		return
	}

	//respond
	c.JSON(http.StatusOK, gin.H{
        "message": "success",
        "created": user,
    })
}