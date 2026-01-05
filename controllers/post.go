package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jahleeljackson/go-rest-api/models"
)

type CreateLogInput struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

func CreateLog(c *gin.Context) {
	var input CreateLogInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post := models.Log{Title: input.Title, Content: input.Content}
	models.DB.Create(&post)

	c.JSON(http.StatusOK, gin.H{"data": post})
}

func FindLogs(c *gin.Context) {
	var logs []models.Log
	models.DB.Find(&logs)

	c.JSON(http.StatusOK, gin.H{"data": logs})
}

func FindLog(c *gin.Context) {
	var post models.Log

	if err := models.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": post})
}

type UpdateLogInput struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Path    string `json:"file_path"`
}

func UpdateLog(c *gin.Context) {
	var log models.Log
	if err := models.DB.Where("id = ?", c.Param("id")).First(&log).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
	}

	var input UpdateLogInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedLog := models.Log{Title: input.Title, Content: input.Content}

	models.DB.Model(&log).Updates(&updatedLog)
	c.JSON(http.StatusOK, gin.H{"data": log})
}

func DeletePost(c *gin.Context) {
	var log models.Log
	if err := models.DB.Where("id = ?", c.Param("id")).First(&log).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
	}

	models.DB.Delete(&log)
	c.JSON(http.StatusOK, gin.H{"data": "success"})
}
