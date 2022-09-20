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
	teams.GET("/", getTeamsByLeague)
	teams.GET("/:teamId", getTeamById)
	teams.DELETE("/:teamId", deleteTeamById)

	r.Run()
}
