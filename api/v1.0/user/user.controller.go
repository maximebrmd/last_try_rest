package user

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"last_try_rest/models"
	"last_try_rest/repository"
	"net/http"
)

func updateAvatar(c *gin.Context) {
	userForm := &models.UserForm{}

	id := c.Params.ByName("id")
	if err := c.ShouldBind(&userForm); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if userForm.Avatar == nil {
		c.JSON(http.StatusBadRequest, errors.New("Avatar should be binded"))
		return
	}

	userID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	user, err := repository.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	file, err := userForm.Avatar.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	user.Avatar, err = models.EncodeFile(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	err = repository.UpdateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, true)
}

func updateUser(c *gin.Context) {
	userForm := &models.UserForm{}

	id := c.Params.ByName("id")
	if err := c.ShouldBind(&userForm); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	userID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	user, err := repository.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	if userForm.Username != nil {
		user.Username = userForm.Username
	}
	if userForm.Stance != nil {
		user.Stance = userForm.Stance
	}

	err = repository.UpdateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, true)
}
