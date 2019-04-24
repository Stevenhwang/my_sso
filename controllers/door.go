package controllers

import (
	"io"
	"net/http"
	"niubility_sso/models"
	"os"
	"strconv"

	echo "github.com/labstack/echo/v4"
)

var pwd, _ = os.Getwd()
var imgDir = pwd + string(os.PathSeparator) + "imgs" + string(os.PathSeparator)

func init() {
	if _, err := os.Stat(imgDir); os.IsNotExist(err) {
		os.Mkdir(imgDir, os.ModePerm)
	}
}

func ListDoors(c echo.Context) error {
	doors := new([]models.Door)
	models.DB.Find(&doors)
	return c.JSON(http.StatusOK, doors)
}

func FindDoor(c echo.Context) error {
	door := new(models.Door)
	uid, _ := strconv.Atoi(c.Param("id"))
	models.DB.First(&door, uid)
	return c.JSON(http.StatusOK, door)
}

func CreateDoor(c echo.Context) error {
	door := new(models.Door)
	if err := c.Bind(door); err != nil {
		return err
	}

	img, err := c.FormFile("img")
	if err != nil {
		return err
	}
	src, err := img.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(imgDir + img.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return err
	}
	// FIX ME!
	door.Img = "https://sso.itcom888.com/api/imgs/" + img.Filename
	models.DB.Create(&door)
	return c.JSON(http.StatusOK, door)
}

func UpdateDoor(c echo.Context) error {
	newDoor := new(models.Door)
	oldDoor := new(models.Door)
	if err := c.Bind(newDoor); err != nil {
		return err
	}
	uid, _ := strconv.Atoi(c.Param("id"))
	models.DB.First(&oldDoor, uid)
	if oldDoor.Name != newDoor.Name && newDoor.Name != "" {
		oldDoor.Name = newDoor.Name
	}
	if oldDoor.Operator != newDoor.Operator && newDoor.Operator != "" {
		oldDoor.Operator = newDoor.Operator
	}
	if oldDoor.Status != newDoor.Status {
		oldDoor.Status = newDoor.Status
	}

	if img, err := c.FormFile("img"); err == nil {
		src, err := img.Open()
		if err != nil {
			return err
		}
		defer src.Close()
		dst, err := os.Create(imgDir + img.Filename)
		if err != nil {
			return err
		}
		defer dst.Close()
		if _, err = io.Copy(dst, src); err != nil {
			return err
		}
		// FIX ME
		oldDoor.Img = "https://sso.itcom888.com/api/imgs/" + img.Filename
	}

	if oldDoor.Link != newDoor.Link && newDoor.Link != "" {
		oldDoor.Link = newDoor.Link
	}
	models.DB.Save(&oldDoor)
	return c.JSON(http.StatusOK, oldDoor)
}

func DeleteDoor(c echo.Context) error {
	door := new(models.Door)
	uid, _ := strconv.Atoi(c.Param("id"))
	models.DB.First(&door, uid)
	models.DB.Delete(&door)
	return c.JSON(http.StatusOK, "Delete done!")
}

func Img(c echo.Context) error {
	img := c.Param("img")
	return c.File(imgDir + img)
}
