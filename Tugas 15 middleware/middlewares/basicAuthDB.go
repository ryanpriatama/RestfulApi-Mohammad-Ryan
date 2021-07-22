package middlewares

import (
	"myapp/project/config"
	"myapp/project/models"

	"github.com/labstack/echo"
)

func BasicAuthDb(username, password string, c echo.Context) (bool, error) {
	db := config.DB
	var user models.User
	if err := db.Where("email=? and password=?", username, password).First(&user).Error; err != nil {
		return false, nil
	}
	return true, nil
}
