package main

import (
	"connorlucier/match-picker/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getLeagues(c *gin.Context) {
	page, pageSize, offset := getPaginationParams(c)
	var leagues []model.League
	var totalRecords int64

	db.Limit(pageSize).Offset(offset).Order("name, created_at DESC").Find(&leagues)
	db.Model(&model.League{}).Count(&totalRecords)

	c.JSON(http.StatusOK, model.Paginated[model.League]{
		Page:         page,
		PageSize:     pageSize,
		TotalRecords: totalRecords,
		Data:         leagues,
	})
}

func getLeagueById(c *gin.Context) {
	var league model.League
	id := c.Param("leagueId")
	db.Limit(1).Preload("Teams").Preload("Seasons").Find(&league, "id = ?", id)

	if league.ID == nil {
		c.Status(http.StatusNoContent)
		return
	}

	c.JSON(http.StatusOK, league)
}

func createLeague(c *gin.Context) {
	var league model.League
	readBody(c, &league)

	if league.Teams == nil {
		league.Teams = []model.Team{}
	}

	db.Create(&league)

	if league.ID == nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusCreated, league)
}

func deleteLeagueById(c *gin.Context) {
	id := c.Param("leagueId")
	db.Select("Teams").Delete(&model.League{ID: &id})
	c.Status(http.StatusNoContent)
}
