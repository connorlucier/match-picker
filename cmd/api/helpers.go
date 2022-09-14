package main

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getDB() *gorm.DB {
	db, err := gorm.Open(postgres.Open("host=localhost user=mpadmin password=mpadmin dbname=matchpicker port=25432 sslmode=disable timezone=UTC"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}

func readBody(c *gin.Context, dest interface{}) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.Status(http.StatusBadRequest)
		panic(err)
	}

	err = json.Unmarshal(body, &dest)
	if err != nil {
		c.Status(http.StatusBadRequest)
		panic(err)
	}
}
