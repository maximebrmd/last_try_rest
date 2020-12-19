package trickTips

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"last_try_rest/models"
	"last_try_rest/repository"
	"net/http"
)

func createTrickTips(c *gin.Context) {
	trickTips := &models.TrickTips{}

	if err := c.ShouldBindJSON(trickTips); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	oid, err := repository.CreateTrickTips(trickTips)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	trickTips.ID = *oid

	c.JSON(http.StatusOK, trickTips)
}

func getAllTrickTips(c *gin.Context) {
	query := &models.Query{
		Filters: map[string]interface{}{},
		Sort:    map[string]interface{}{},
	}

	filters := c.Query("filters")
	if len(filters) <= 0 {
		filters = "{}"
	}
	err := json.Unmarshal([]byte(filters), &query.Filters)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	sort := c.Query("sort")
	if len(sort) <= 0 {
		sort = "{}"
	}
	err = json.Unmarshal([]byte(sort), &query.Sort)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	fmt.Println(query.Filters)
	fmt.Println(query.Sort)

	allTrickTips, err := repository.GetAllTrickTips(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, allTrickTips)
}

func updateTrickTipsImages(c *gin.Context) {
	trickTipsForm := &models.TrickTipsForm{}

	id := c.Params.ByName("id")
	if err := c.ShouldBind(&trickTipsForm); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if trickTipsForm.Sequence == nil || trickTipsForm.Thumbnail == nil {
		c.JSON(http.StatusBadRequest, errors.New("Sequence and Thumbnail should be binded"))
		return
	}

	trickTipsID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	trickTips, err := repository.GetTrickTipsByID(trickTipsID)
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	// Encode Thumbnail
	file, err := trickTipsForm.Thumbnail.Open()
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
	trickTips.Sequence = make([]*string, 0)
	for _, sequence := range trickTipsForm.Sequence {
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

	err = repository.UpdateTrickTips(trickTips)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, true)
}
