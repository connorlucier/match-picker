package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const default_page_size = 100
const max_page_size = 1000

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

func getPaginationParams(c *gin.Context) (int, int, int) {
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))

	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = default_page_size
	}

	if pageSize > max_page_size {
		pageSize = max_page_size
	}

	return page, pageSize, (page - 1) * pageSize
}
