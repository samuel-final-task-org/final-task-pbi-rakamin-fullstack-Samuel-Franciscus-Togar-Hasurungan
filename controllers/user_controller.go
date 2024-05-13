package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"final-task/app"
	"final-task/helpers"
)

// UserController adalah controller untuk entitas User.
type UserController struct {
	DB *gorm.DB
}

// registerUser digunakan untuk mendaftarkan pengguna baru.
func (ctrl UserController) registerUser(c *gin.Context) {
	var user app.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak valid"})
		return
	}

	hashedPassword, _ := helpers.HashPassword(user.Password)
	user.Password = hashedPassword

	if result := ctrl.DB.Create(&user); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Pengguna berhasil didaftarkan", "user": user})
}

// loginUser digunakan untuk proses autentikasi pengguna.
func (ctrl UserController) loginUser(c *gin.Context) {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak valid"})
		return
	}

	var user app.User
	result := ctrl.DB.Where("email = ?", credentials.Email).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email atau password tidak valid"})
		return
	}

	if !helpers.CheckPasswordHash(credentials.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email atau password tidak valid"})
		return
	}

	token, err := helpers.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghasilkan token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login berhasil", "token": token})
}

// updateUser digunakan untuk memperbarui data pengguna.
func (ctrl UserController) updateUser(c *gin.Context) {
	var userUpdates app.User
	userID := c.Param("id")

	if err := c.ShouldBindJSON(&userUpdates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak valid"})
		return
	}

	result := ctrl.DB.Model(&app.User{}).Where("id = ?", userID).Updates(userUpdates)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pengguna berhasil diperbarui"})
}

// deleteuser digunakan untuk menghapus data pengguna.
func (ctrl UserController) deleteuser(c *gin.Context) {
	userID := c.Param("id")

	result := ctrl.DB.Delete(&app.User{}, userID)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tidak ada pengguna yang ditemukan dengan ID yang diberikan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pengguna berhasil dihapus"})
}
