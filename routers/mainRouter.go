package routers

import (
	"database/sql"
	"quiz3/controllers"
	"quiz3/middlewares"

	"github.com/gin-gonic/gin"
)

func StartServer(db *sql.DB) *gin.Engine {
	router := gin.Default()

	categoryRoutes := router.Group("/api/categories")
	categoryRoutes.Use(middlewares.BasicAuth(db))
	{
		categoryRoutes.GET("", func(c *gin.Context) {
			controllers.GetCategories(c, db)
		})

		categoryRoutes.GET("/:id", func(c *gin.Context) {
			controllers.GetCategoryByID(c, db)
		})

		categoryRoutes.POST("", func(c *gin.Context) {
			controllers.CreateCategory(c, db)
		})

		categoryRoutes.PUT("/:id", func(c *gin.Context) {
			controllers.UpdateCategory(c, db)
		})

		categoryRoutes.DELETE("/:id", func(c *gin.Context) {
			controllers.DeleteCategory(c, db)
		})

		categoryRoutes.GET("/:id/books", func(c *gin.Context) {
			controllers.GetBooksByCategoryID(c, db)
		})
	}

	bookRoutes := router.Group("/api/books")
	bookRoutes.Use(middlewares.BasicAuth(db))
	{
		bookRoutes.GET("", func(c *gin.Context) {
			controllers.GetBooks(c, db)
		})
		bookRoutes.GET("/:id", func(c *gin.Context) {
			controllers.GetBookByID(c, db)
		})
		bookRoutes.POST("", func(c *gin.Context) {
			controllers.CreateBook(c, db)
		})
		bookRoutes.PUT("/:id", func(c *gin.Context) {
			controllers.UpdateBook(c, db)
		})
		bookRoutes.DELETE("/:id", func(c *gin.Context) {
			controllers.DeleteBook(c, db)
		})
	}
	return router
}
