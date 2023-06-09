package main

import (
	"connorlucier/match-picker/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getPicks(c *gin.Context) {
	userId := c.Param("userId")
	page, pageSize, offset := getPaginationParams(c)

	var picks []model.Pick
	var totalRecords int64

	db.Limit(pageSize).Offset(offset).Find(&picks, "user_id = ?", userId)
	db.Model(&model.Pick{}).Where("user_id = ?", userId).Count(&totalRecords)

	c.JSON(http.StatusOK, model.Paginated[model.Pick]{
		Page:         page,
		PageSize:     pageSize,
		TotalRecords: totalRecords,
		Data:         picks,
	})
}

func getPickById(c *gin.Context) {
	var pick model.Pick
	id := c.Param("pickId")
	db.Limit(1).Find(&pick, "id = ?", id)

	if pick.ID == nil {
		c.Status(http.StatusNoContent)
		return
	}

	c.JSON(http.StatusOK, pick)
}

func createPick(c *gin.Context) {
	// TODO get user from auth, not body
	var pick model.Pick
	readBody(c, &pick)
	db.Create(&pick)

	if pick.ID == nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusCreated, pick)
}

func updatePick(c *gin.Context) {
	// TODO get user from auth
	pickId := c.Param("pickId")
	var pick model.Pick
	readBody(c, &pick)
	pick.ID = &pickId
	db.Save(&pick)

	c.JSON(http.StatusOK, pick)
}

func deletePick(c *gin.Context) {
	// TODO get user from auth
	id := c.Param("pickId")
	db.Delete(&model.Pick{}, "id = ?", id)
	c.Status(http.StatusNoContent)
}
