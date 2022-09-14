package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func allHealthChecks(c *gin.Context) {
	checkDb()
	c.JSON(http.StatusOK, "Healthy")
}

func dbHealthCheck(c *gin.Context) {
	checkDb()
	c.JSON(http.StatusOK, "Healthy")
}

func checkDb() {
	var ok bool
	db.Raw("select true;").Scan(&ok)
	if db.Error != nil {
		panic(db.Error)
	}

	if !ok {
		panic(errors.New("failed to connect to db"))
	}
}
