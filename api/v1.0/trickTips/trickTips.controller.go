package trickTips

import (
	"last_try_rest/db"
	"last_try_rest/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func addTrickTips(c *gin.Context) {
	trickTips := &models.TrickTips{}

	if err := c.ShouldBindJSON(trickTips); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	oid, err := db.AddTrickTips(trickTips)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, oid)
}

func getAllTrickTips(c *gin.Context) {
	allTrickTips, err := db.GetAllTrickTips()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, allTrickTips)
}
