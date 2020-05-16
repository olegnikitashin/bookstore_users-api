package users

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/olegnikitashin/bookstore_users-api/domains/users"
	"github.com/olegnikitashin/bookstore_users-api/services"
	"github.com/olegnikitashin/bookstore_users-api/utils/errors"
)

func CreateUser(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid JSON body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	fmt.Println(user)
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}
