package router

import (
	"net/http"
	"niubility_sso/controllers"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// initialize new server
var Server = echo.New()

func SetRouter() {
	// middleware
	Server.Use(middleware.Logger())
	Server.Use(middleware.Recover())
	Server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))
	// index
	api := Server.Group("/api")
	rapi := Server.Group("/rapi")
	rapi.Use(middleware.JWT([]byte("secret")))
	api.GET("/", controllers.Welcome)
	// user
	api.GET("/users", controllers.ListUsers)
	api.GET("/user/:id", controllers.FindUser)
	api.POST("/users", controllers.CreateUser)
	api.PUT("/user/:id", controllers.UpdateUser)
	api.DELETE("/user/:id", controllers.DeleteUser)
	api.POST("/user/login", controllers.Login)
	rapi.POST("/user/logout", controllers.Logout)
	// door
	api.GET("/doors", controllers.ListDoors)
	api.GET("/door/:id", controllers.FindDoor)
	api.POST("/doors", controllers.CreateDoor)
	api.PUT("/door/:id", controllers.UpdateDoor)
	api.DELETE("/door/:id", controllers.DeleteDoor)
	api.GET("/imgs/:img", controllers.Img)
}
