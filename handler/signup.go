package handler

import (
	"awesomeProject/controller"
	"awesomeProject/gin_auth/globals"
	"awesomeProject/gin_auth/helpers"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SignupGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		if user != nil {
			c.HTML(http.StatusBadRequest, "dashboard.html",
				gin.H{
					"content": "Please logout first",
					"user":    user,
				})
			return
		}
		c.HTML(http.StatusOK, "signup.html", gin.H{
			"content": "",
			"user":    user,
		})
	}
}

func SignupPostHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)

		email := c.PostForm("email")
		username := c.PostForm("username")
		password := c.PostForm("password")

		if helpers.EmptyUserPassSignUp(username, password, email) {
			c.HTML(http.StatusBadRequest, "signup.html", gin.H{"content": "Parameters can't be empty"})
			return
		}

		if _, err := controller.FindByField("email", email); err == nil {
			c.HTML(http.StatusBadRequest, "signup.html", gin.H{"content": "Account alredy exists, please login"})
			return
		}

		if c.PostForm("password") == c.PostForm("passwordAgain") {
			err := controller.SaveUser(username, password, email)
			if err != nil {
				log.Printf("Error while saving user")
			}
			c.HTML(http.StatusOK, "signup.html", gin.H{"content": "User has been created successfully"})
		} else {
			c.HTML(http.StatusBadRequest, "signup.html", gin.H{"content": "Passwords don't match"})
			return
		}
		session.Set(globals.Userkey, username)
		if err := session.Save(); err != nil {
			c.HTML(http.StatusInternalServerError, "signup.html", gin.H{"content": "Failed to save session"})
			return
		}
		c.Redirect(http.StatusFound, "/login")
	}
}
