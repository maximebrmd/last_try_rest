package trickTips

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"last_try_rest/db"
	"last_try_rest/models"
	"net/http"
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
	trickTips.ID = *oid

	c.JSON(http.StatusOK, trickTips)
}

func addTrickTipsImages(c *gin.Context) {
	trickTipsImages := &models.TrickTipsImages{}
	trickTips := &models.TrickTips{}

	if err := c.ShouldBind(&trickTipsImages); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// Encode Thumbnail
	file, err := trickTipsImages.Thumbnail.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	trickTips.Thumbnail, err = models.EncodeFile(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	// Encode Sequence
	for _, sequence := range trickTipsImages.Sequence {
		file, err := sequence.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		base64, err := models.EncodeFile(file)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		trickTips.Sequence = append(trickTips.Sequence, base64)
	}

	fmt.Println(*trickTipsImages.TrickTipsID)

	trickTipsID, err := primitive.ObjectIDFromHex(*trickTipsImages.TrickTipsID)
	if err != nil {
		panic(err)
	}

	fmt.Println(trickTipsID)

	err = db.AddTrickTipsImages(trickTips, trickTipsID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, true)
}

func getAllTrickTips(c *gin.Context) {
	allTrickTips, err := db.GetAllTrickTips()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, allTrickTips)
}
