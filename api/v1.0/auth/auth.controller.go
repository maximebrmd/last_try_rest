package auth

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"last_try_rest/models"
	"last_try_rest/repository"
	"net/http"
)

func createUser(c *gin.Context) {
	user := &models.User{}

	if err := c.ShouldBindJSON(user); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	query := &models.Query{
		Filters: map[string]interface{}{
			"email": user.Email,
		},
	}

	users, err := repository.GetAllUser(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if len(users) > 0 {
		c.JSON(http.StatusConflict, errors.New("email already exists"))
		return
	}

	//JWT & Hash PWD

	oid, err := repository.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	user.ID = *oid

	c.JSON(http.StatusOK, user)
}

func loginUser(c *gin.Context) {
	user := &models.User{}

	if err := c.ShouldBindJSON(user); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	users, err := repository.GetAllUser(nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	for _, user := range users {
		// TODO: Check User exist & credentials are ok || //JWT & Hash PWD
		fmt.Println(user)
	}

	c.JSON(http.StatusOK, user)
}
