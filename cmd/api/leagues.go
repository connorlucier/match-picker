package main

import (
	"connorlucier/match-picker/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getLeagues(c *gin.Context) {
	// TODO pagination
	var leagues []model.League
	db.Find(&leagues)
	c.JSON(http.StatusOK, leagues)
}

func getLeagueById(c *gin.Context) {
	var league model.League
	id := c.Param("leagueId")
	db.Limit(1).Preload("Teams").Find(&league, "id = ?", id)

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
