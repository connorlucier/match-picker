package main

import (
	"connorlucier/match-picker/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getGroups(c *gin.Context) {
	page, pageSize, offset := getPaginationParams(c)

	var groups []model.Group
	var totalRecords int64

	db.Limit(pageSize).Offset(offset).Find(&groups)
	db.Model(&model.Group{}).Count(&totalRecords)

	c.JSON(http.StatusOK, model.Paginated[model.Group]{
		Page:         page,
		PageSize:     pageSize,
		TotalRecords: totalRecords,
		Data:         groups,
	})
}

func getGroupById(c *gin.Context) {
	var group model.Group
	id := c.Param("groupId")
	db.Limit(1).Find(&group, "id = ?", id)

	if group.ID == nil {
		c.Status(http.StatusNoContent)
		return
	}

	c.JSON(http.StatusOK, group)
}

func createGroup(c *gin.Context) {
	var group model.Group
	readBody(c, &group)
	db.Create(&group)

	if group.ID == nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusCreated, group)
}

func updateGroup(c *gin.Context) {
	id := c.Param("groupId")
	var group model.Group
	readBody(c, &group)
	group.ID = &id
	db.Save(&group)

	c.JSON(http.StatusOK, group)
}

func deleteGroup(c *gin.Context) {
	id := c.Param("groupId")
	db.Delete(&model.Group{}, "id = ?", id)
	c.Status(http.StatusNoContent)
}
