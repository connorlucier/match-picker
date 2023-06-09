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

	users := r.Group("/users")
	users.POST("/", createUser)
	users.GET("/", getUsers)
	users.GET("/:userId", getUserById)
	users.PUT("/:userId", updateUser)
	users.DELETE("/:userId", deleteUser)

	picks := users.Group("/:userId/picks")
	picks.POST("/", createPick)
	picks.GET("/", getPicks)
	picks.GET("/:pickId", getPickById)
	picks.PUT("/:pickId", updatePick)
	picks.DELETE("/:pickId", deletePick)

	groups := r.Group("/groups")
	groups.POST("/", createGroup)
	groups.GET("/", getGroups)
	groups.GET("/:groupId", getGroupById)
	groups.PUT("/:groupId", updateGroup)
	groups.DELETE("/:groupId", deleteGroup)

	leagues := r.Group("/leagues")
	leagues.POST("/", createLeague)
	leagues.GET("/", getLeagues)
	leagues.GET("/:leagueId", getLeagueById)
	leagues.PUT("/:leagueId", updateLeague)
	leagues.DELETE("/:leagueId", deleteLeague)

	teams := leagues.Group("/:leagueId/teams")
	teams.POST("/", createTeam)
	teams.GET("/", getTeams)
	teams.GET("/:teamId", getTeamById)
	teams.PUT("/:teamId", updateTeam)
	teams.DELETE("/:teamId", deleteTeamById)

	seasons := leagues.Group("/:leagueId/seasons")
	seasons.POST("/", createSeason)
	seasons.GET("/", getSeasons)
	seasons.GET("/:seasonId", getSeasonById)
	seasons.PUT("/:seasonId", updateSeason)
	seasons.DELETE("/:seasonId", deleteSeason)

	matches := seasons.Group("/:seasonId/matches")
	matches.POST("/", createMatch)
	matches.GET("/", getMatches)
	matches.GET("/:matchId", getMatchById)
	matches.PUT("/:matchId", updateMatch)
	matches.DELETE("/:matchId", deleteMatch)

	r.Run()
}
