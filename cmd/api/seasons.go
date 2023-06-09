package main

import (
	"connorlucier/match-picker/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getSeasons(c *gin.Context) {
	leagueId := c.Param("leagueId")
	page, pageSize, offset := getPaginationParams(c)

	var seasons []model.Season
	var totalRecords int64

	db.Limit(pageSize).Offset(offset).Order("start_year DESC, created_at DESC").Find(&seasons, "league_id = ?", leagueId)
	db.Model(&model.Season{}).Where("league_id = ?", leagueId).Count(&totalRecords)

	c.JSON(http.StatusOK, model.Paginated[model.Season]{
		Page:         page,
		PageSize:     pageSize,
		TotalRecords: totalRecords,
		Data:         seasons,
	})
}

func getSeasonById(c *gin.Context) {
	var season model.Season
	id := c.Param("seasonId")
	db.Limit(1).Preload("Matches").Find(&season, "id = ?", id)

	if season.ID == nil {
		c.Status(http.StatusNoContent)
		return
	}

	c.JSON(http.StatusOK, season)
}

func createSeason(c *gin.Context) {
	var season model.Season
	readBody(c, &season)
	db.Create(&season)

	if season.ID == nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusCreated, season)
}

func updateSeason(c *gin.Context) {
	id := c.Param("seasonId")
	var season model.Season
	readBody(c, &season)
	season.ID = &id
	db.Save(&season)

	c.JSON(http.StatusOK, season)
}

func deleteSeason(c *gin.Context) {
	id := c.Param("seasonId")
	db.Delete(&model.Season{}, "id = ?", id)
	c.Status(http.StatusNoContent)
}
