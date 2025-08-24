package routers

import (
	"database/sql"

	"Quiz/controllers"
	"Quiz/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(db *sql.DB) *gin.Engine {
	r := gin.Default()

	// Controller instances
	bookCtrl := controllers.BookController{DB: db}
	catCtrl := controllers.CategoryController{DB: db}

	// Public routes (Auth)
	r.POST("/api/users/register", controllers.Register)
	r.POST("/api/users/login", controllers.Login)

	// Protected routes (JWT required)
	api := r.Group("/api", middleware.JWTAuthMiddleware())
	{
		// Books
		api.POST("/books", bookCtrl.CreateBook)
		api.GET("/books", bookCtrl.GetAllBooks)
		api.GET("/books/:id", bookCtrl.GetBookByID)
		api.PUT("/books/:id", bookCtrl.UpdateBook)
		api.DELETE("/books/:id", bookCtrl.DeleteBook)

		// Categories
		api.POST("/categories", catCtrl.CreateCategory)
		api.GET("/categories", catCtrl.GetAllCategories)
		api.GET("/categories/:id", catCtrl.GetCategoryByID)
		api.DELETE("/categories/:id", catCtrl.DeleteCategory)
		api.GET("/categories/:id/books", catCtrl.GetBooksByCategory)
	}

	return r
}
