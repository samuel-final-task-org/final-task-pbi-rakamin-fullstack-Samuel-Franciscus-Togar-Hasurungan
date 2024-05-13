package router

import (
	"github.com/gin-gonic/gin"
	"final-task/controllers"
	"final-task/database"
	"final-task/middlewares"
)

// setupRouter mengatur rute-rute aplikasi dan mengembalikan instance gin.Engine.
func setupRouter() *gin.Engine {
	r := gin.Default()

	// Inisialisasi koneksi database
	db := database.InitDB()

	// Inisialisasi controller
	userController := controllers.UserController{DB: db}
	photoController := controllers.PhotoController{DB: db}

	// Rute-rute untuk pengguna
	r.POST("/users/register", userController.registerUser)
	r.POST("/users/login", userController.loginUser)
	r.PUT("/users/:id", userController.updateUser)
	r.DELETE("/users/:id", userController.deleteuser)

	// Grup rute untuk foto dengan middleware otentikasi
	photoRoutes := r.Group("/photos")
	photoRoutes.Use(middlewares.AuthMiddleware())
	{
		photoRoutes.GET("/", photoController.getPhotos)
		photoRoutes.POST("/", photoController.uploadPhoto)
		photoRoutes.PUT("/:id", photoController.updatePhoto)
		photoRoutes.DELETE("/:id", photoController.deletePhoto)
	}

	return r
}

