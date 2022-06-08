package routes

import (
	"awesomeProject/handler"

	"github.com/gin-gonic/gin"
)

func PublicRoutes(g *gin.RouterGroup) {

	g.GET("/login", handler.LoginGetHandler())
	g.POST("/login", handler.LoginPostHandler())
	g.GET("/", handler.IndexGetHandler())
	g.GET("/signup", handler.SignupGetHandler())
	g.POST("/signup", handler.SignupPostHandler())

}

func PrivateRoutes(g *gin.RouterGroup) {
	g.GET("/dashboard", handler.DashboardGetHandler())
	g.GET("/logout", handler.LogoutGetHandler())
}
