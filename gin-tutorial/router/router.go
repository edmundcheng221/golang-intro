package router

import (
	handler "EdmundsBankai/golang-intro/gin-tutorial/handlers"

	"github.com/gin-gonic/gin"
)

func v1RouteGrouping(router *gin.Engine) *gin.RouterGroup {
	api := router.Group("/api")
	v1 := api.Group("/v1")
	return v1
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	grouping := v1RouteGrouping(router)
	grouping.GET("/albums", handler.GetAlbums)
	grouping.POST("/albums", handler.PostAlbum)
	grouping.PUT("/albums", handler.UpdateAlbum)
	grouping.DELETE("/albums", func(c *gin.Context) {
		id := c.Query("id")
		handler.DeleteAlbum(c, id)
	})
	grouping.PATCH("/albums", func(c *gin.Context) {
		handler.PatchAlbum(c)
	})
	return router
}
