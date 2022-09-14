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

func getLeague(c *gin.Context) {
	var league model.League
	id := c.Param("leagueId")
	db.Limit(1).Find(&league, "id = ?", id)

	if league.ID == nil {
		c.Status(http.StatusNoContent)
		return
	}

	c.JSON(http.StatusOK, league)
}

func createLeague(c *gin.Context) {
	var league model.League
	readBody(c, &league)
	db.Create(&league)

	if league.ID == nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusCreated, league)
}
