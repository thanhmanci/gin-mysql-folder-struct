package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"newfeed/db"
	"newfeed/models"

	"github.com/gin-gonic/gin"
)

var database = db.InitDatabase()

// CreateNewfeed blablaba
func CreateNewfeed() gin.HandlerFunc {
	return func(c *gin.Context) {
		completed, _ := strconv.Atoi(c.PostForm("completed"))
		newfeed := &models.NewfeedModel{Title: c.PostForm("title"), Completed: completed}
		fmt.Println(database)
		database.Save(&newfeed)
		c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "newfeed item created successfully!", "resourceId": newfeed.Title})
	}
}

// FetchAllNewfeed blablaba
func FetchAllNewfeed() gin.HandlerFunc {
	return func(c *gin.Context) {
		var newfeeds []models.NewfeedModel
		var _Newfeeds []models.TransformedNewfeed
		database.Find(&newfeeds)
		if len(newfeeds) <= 0 {
			c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No newfeed found!"})
			return
		}
		//transforms the newfeeds for building a good response
		for _, item := range newfeeds {
			completed := false
			if item.Completed == 1 {
				completed = true
			} else {
				completed = false
			}
			_Newfeeds = append(_Newfeeds, models.TransformedNewfeed{ID: item.ID, Title: item.Title, Completed: completed})
		}
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _Newfeeds})
	}
}

// FetchSingleNewfeed blablaba
func FetchSingleNewfeed() gin.HandlerFunc {
	return func(c *gin.Context) {
		var newfeed models.NewfeedModel
		newfeedID := c.Param("id")
		fmt.Println(newfeedID)
		database.First(&newfeed, newfeedID)
		if newfeed.ID == 0 {
			c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No newfeed found!"})
			return
		}
		completed := false
		if newfeed.Completed == 1 {
			completed = true
		} else {
			completed = false
		}
		_Newfeed := models.TransformedNewfeed{ID: newfeed.ID, Title: newfeed.Title, Completed: completed}
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _Newfeed})
	}
}

// UpdateNewfeed update a newfeed
func UpdateNewfeed() gin.HandlerFunc {
	return func(c *gin.Context) {
		var newfeed models.NewfeedModel
		newfeedID := c.Param("id")
		database.First(&newfeed, newfeedID)
		if newfeed.ID == 0 {
			c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No newfeed found!"})
			return
		}
		database.Model(&newfeed).Update("title", c.PostForm("title"))
		completed, _ := strconv.Atoi(c.PostForm("completed"))
		database.Model(&newfeed).Update("completed", completed)
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "newfeed updated successfully!"})
	}
}

// DeleteNewfeed remove a newfeed
func DeleteNewfeed() gin.HandlerFunc {
	return func(c *gin.Context) {
		var newfeed models.NewfeedModel
		newfeedID := c.Param("id")
		database.First(&newfeed, newfeedID)
		if newfeed.ID == 0 {
			c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No newfeed found!"})
			return
		}
		database.Delete(&newfeed)
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "newfeed deleted successfully!"})
	}
}
