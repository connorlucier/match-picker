package main

import (
	"connorlucier/match-picker/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getMatches(c *gin.Context) {
	seasonId := c.Param("seasonId")
	after, before := getDateRangeParams(c)
	page, pageSize, offset := getPaginationParams(c)

	var matches []model.Match
	var totalRecords int64

	db.Limit(pageSize).Offset(offset).Order("starts_at ASC").
		Where("? OR starts_at >= ?", after == nil, after).
		Where("? OR starts_at <= ?", before == nil, before).
		Find(&matches, "season_id = ?", seasonId)
	db.Model(&model.Match{}).
		Where("season_id = ?", seasonId).
		Where("? OR starts_at >= ?", after == nil, after).
		Where("? OR starts_at <= ?", before == nil, before).
		Count(&totalRecords)

	c.JSON(http.StatusOK, model.Paginated[model.Match]{
		Page:         page,
		PageSize:     pageSize,
		TotalRecords: totalRecords,
		Data:         matches,
	})
}

func getMatchById(c *gin.Context) {
	var match model.Match
	id := c.Param("matchId")
	db.Limit(1).Find(&match, "id = ?", id)

	if match.ID == nil {
		c.Status(http.StatusNoContent)
		return
	}

	c.JSON(http.StatusOK, match)
}

func createMatch(c *gin.Context) {
	seasonId := c.Param("seasonId")
	var match model.Match
	match.SeasonID = &seasonId
	readBody(c, &match)
	db.Create(&match)

	if match.ID == nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusCreated, match)
}

func updateMatch(c *gin.Context) {
	id := c.Param("matchId")
	var match model.Match
	readBody(c, &match)
	match.ID = &id
	db.Save(&match)

	c.JSON(http.StatusOK, match)
}

func deleteMatch(c *gin.Context) {
	id := c.Param("matchId")
	db.Delete(&model.Match{}, "id = ?", id)
	c.Status(http.StatusNoContent)
}
