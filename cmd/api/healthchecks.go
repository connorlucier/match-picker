package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func allHealthChecks(c *gin.Context) {
	checkDb(c)
	healthy(c)
}

func dbHealthCheck(c *gin.Context) {
	checkDb(c)
	healthy(c)
}

func checkDb(c *gin.Context) {
	var ok bool
	db.Raw("select true;").Scan(&ok)
	if db.Error != nil {
		unhealthy(c)
		panic(db.Error)
	}

	if !ok {
		unhealthy(c)
		panic(errors.New("failed to connect to db"))
	}
}

func healthy(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{
		"status": "healthy",
	})
}

func unhealthy(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, map[string]string{
		"status": "unhealthy",
	})
}
