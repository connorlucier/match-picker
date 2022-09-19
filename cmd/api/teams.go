package main

import (
	"connorlucier/match-picker/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getTeamsByLeague(c *gin.Context) {
	id := c.Param("leagueId")

	// TODO pagination
	var teams []model.Team
	db.Find(&teams, "league_id = ?", id)
	c.JSON(http.StatusOK, teams)
}

func getTeamById(c *gin.Context) {
	var team model.Team
	id := c.Param("teamId")
	db.Limit(1).Find(&team, "id = ?", id)

	if team.ID == nil {
		c.Status(http.StatusNoContent)
		return
	}

	c.JSON(http.StatusOK, &team)
}

func createTeam(c *gin.Context) {
	var team model.Team
	readBody(c, &team)
	db.Create(&team)

	if team.ID == nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusCreated, team)
}

func deleteTeamById(c *gin.Context) {
	id := c.Param("teamId")
	db.Delete(&model.Team{}, "id = ?", id)
	c.Status(http.StatusNoContent)
}
