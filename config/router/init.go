package router

import (
	"newfeed/controllers"

	"github.com/gin-gonic/gin"
)

//Init blababa
func Init() {
	router := gin.Default()
	v1 := router.Group("/api/v1/newfeeds")
	{
		v1.POST("/", controllers.CreateNewfeed())
		v1.GET("/", controllers.FetchAllNewfeed())
		v1.GET("/:id", controllers.FetchSingleNewfeed())
		v1.PUT("/:id", controllers.UpdateNewfeed())
		v1.DELETE("/:id", controllers.DeleteNewfeed())
	}
	router.Run()
}
