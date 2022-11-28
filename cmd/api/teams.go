package main

import (
	"connorlucier/match-picker/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getTeamsByLeagueId(c *gin.Context) {
	leagueId := c.Param("leagueId")
	page, pageSize, offset := getPaginationParams(c)

	var teams []model.Team
	var totalRecords int64

	db.Limit(pageSize).Offset(offset).Order("name, created_at DESC").Find(&teams, "league_id = ?", leagueId)
	db.Model(&model.Season{}).Where("league_id = ?", leagueId).Count(&totalRecords)

	c.JSON(http.StatusOK, model.Paginated[model.Team]{
		Page:         page,
		PageSize:     pageSize,
		TotalRecords: totalRecords,
		Data:         teams,
	})
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
	leagueId := c.Param("leagueId")
	var team model.Team
	team.LeagueID = &leagueId
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
