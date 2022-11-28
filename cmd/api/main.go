package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB = getDB()

func main() {
	r := gin.Default()

	health := r.Group("/health")
	health.GET("/", allHealthChecks)
	health.GET("/db", dbHealthCheck)

	leagues := r.Group("/leagues")
	leagues.POST("/", createLeague)
	leagues.GET("/", getLeagues)
	leagues.GET("/:leagueId", getLeagueById)
	leagues.DELETE("/:leagueId", deleteLeagueById)

	teams := leagues.Group("/:leagueId/teams")
	teams.POST("/", createTeam)
	teams.GET("/", getTeamsByLeagueId)
	teams.GET("/:teamId", getTeamById)
	teams.DELETE("/:teamId", deleteTeamById)

	seasons := leagues.Group("/:leagueId/seasons")
	seasons.POST("/", createSeason)
	seasons.GET("/", getSeasonsByLeagueId)
	seasons.GET("/:seasonId", getSeasonById)
	seasons.DELETE("/:seasonId", deleteSeasonById)

	matches := seasons.Group("/:seasonId/matches")
	matches.POST("/", createMatch)
	matches.GET("/", getMatches)
	matches.GET("/:matchId", getMatchById)
	matches.DELETE("/:matchId", deleteMatchById)

	r.Run()
}
