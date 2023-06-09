package main

import (
	"connorlucier/match-picker/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getUsers(c *gin.Context) {
	page, pageSize, offset := getPaginationParams(c)

	var users []model.User
	var totalRecords int64

	db.Limit(pageSize).Offset(offset).Find(&users)
	db.Model(&model.User{}).Count(&totalRecords)

	c.JSON(http.StatusOK, model.Paginated[model.User]{
		Page:         page,
		PageSize:     pageSize,
		TotalRecords: totalRecords,
		Data:         users,
	})
}

func getUserById(c *gin.Context) {
	var user model.User
	id := c.Param("userId")
	db.Limit(1).Find(&user, "id = ?", id)

	if user.ID == nil {
		c.Status(http.StatusNoContent)
		return
	}

	c.JSON(http.StatusOK, user)
}

func createUser(c *gin.Context) {
	var user model.User
	readBody(c, &user)
	db.Create(&user)

	if user.ID == nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusCreated, user)
}

func updateUser(c *gin.Context) {
	id := c.Param("userId")
	var user model.User
	readBody(c, &user)
	user.ID = &id
	db.Save(&user)

	c.JSON(http.StatusOK, user)
}

func deleteUser(c *gin.Context) {
	id := c.Param("userId")
	db.Delete(&model.User{}, "id = ?", id)
	c.Status(http.StatusNoContent)
}
