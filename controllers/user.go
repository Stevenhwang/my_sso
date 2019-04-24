package controllers

import (
	"net/http"
	"niubility_sso/models"
	"niubility_sso/utils"
	"strconv"

	"time"

	"github.com/dgrijalva/jwt-go"

	echo "github.com/labstack/echo/v4"
)

func ListUsers(c echo.Context) error {
	users := new([]models.User)
	models.DB.Find(&users)
	return c.JSON(http.StatusOK, users)
}

func FindUser(c echo.Context) error {
	user := new(models.User)
	uid, _ := strconv.Atoi(c.Param("id"))
	models.DB.First(&user, uid)
	return c.JSON(http.StatusOK, user)
}

func CreateUser(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return err
	}
	if user.Name == "" || user.Passwd == "" {
		return c.JSON(http.StatusLengthRequired, "Need name or password!")
	}
	user.Passwd = utils.GetEncrypted(user.Passwd)
	models.DB.Create(&user)
	return c.JSON(http.StatusOK, user)
}

func UpdateUser(c echo.Context) error {
	newUser := new(models.User)
	oldUser := new(models.User)
	if err := c.Bind(newUser); err != nil {
		return err
	}
	uid, _ := strconv.Atoi(c.Param("id"))
	models.DB.First(&oldUser, uid)
	if oldUser.Name != newUser.Name && newUser.Name != "" {
		oldUser.Name = newUser.Name
	}
	if oldUser.Passwd != newUser.Passwd && newUser.Passwd != "" {
		oldUser.Passwd = utils.GetEncrypted(newUser.Passwd)
	}
	models.DB.Save(&oldUser)
	return c.JSON(http.StatusOK, oldUser)
}

func DeleteUser(c echo.Context) error {
	user := new(models.User)
	uid, _ := strconv.Atoi(c.Param("id"))
	models.DB.First(&user, uid)
	models.DB.Delete(&user)
	return c.JSON(http.StatusOK, "Delete done!")
}

func Login(c echo.Context) error {
	user := new(models.User)
	name := c.FormValue("name")
	inputPW := c.FormValue("passwd")
	models.DB.Where("name= ?", name).First(&user)
	if user.Name == "" {
		return c.JSON(http.StatusNotFound, "No such user!")
	} else if utils.CheckPW(user.Passwd, inputPW) == "PW Wrong!" {
		return c.JSON(http.StatusUnauthorized, "PW Wrong!")
	} else {
		// Create token
		token := jwt.New(jwt.SigningMethodHS256)
		// Set claims
		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = user.Name
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}
		utils.RefreshToken(name, t)
		return c.JSON(http.StatusOK, map[string]string{
			"token": t,
		})
	}
}

func Logout(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return c.JSON(http.StatusOK, utils.DelToken(name))
}
