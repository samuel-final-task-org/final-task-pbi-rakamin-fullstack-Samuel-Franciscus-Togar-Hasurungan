package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"final-task/app"
	"gorm.io/gorm"
)

// PhotoController adalah controller untuk entitas Photo.
type PhotoController struct {
	DB *gorm.DB
}

// uploadPhoto digunakan untuk mengunggah foto baru.
func (ctrl *PhotoController) uploadPhoto(c *gin.Context) {
	var photo app.Photo
	if err := c.ShouldBind(&photo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak valid"})
		return
	}

	userID, exists := getUserIDFromContext(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Tidak diotorisasi"})
		return
	}
	photo.UserID = userID

	file, err := c.FormFile("photo")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error unggah foto"})
		return
	}
	filePath := "uploads/" + file.Filename
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan foto"})
		return
	}
	photo.PhotoUrl = filePath

	if result := ctrl.DB.Create(&photo); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Foto berhasil diunggah", "photo": photo})
}

// getPhotos digunakan untuk mendapatkan daftar foto milik pengguna yang sedang masuk.
func (ctrl *PhotoController) getPhotos(c *gin.Context) {
	userID, exists := getUserIDFromContext(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Tidak diotorisasi"})
		return
	}

	var photos []app.Photo
	if result := ctrl.DB.Where("user_id = ?", userID).Find(&photos); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"photos": photos})
}

// updatePhoto digunakan untuk memperbarui data foto.
func (ctrl *PhotoController) updatePhoto(c *gin.Context) {
	photoID := c.Param("id")
	var photoUpdates app.Photo
	if err := c.ShouldBindJSON(&photoUpdates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak valid"})
		return
	}

	result := ctrl.DB.Model(&app.Photo{}).Where("id = ?", photoID).Updates(photoUpdates)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Foto berhasil diperbarui"})
}

// deletePhoto digunakan untuk menghapus data foto.
func (ctrl *PhotoController) deletePhoto(c *gin.Context) {
	photoID := c.Param("id")
	result := ctrl.DB.Delete(&app.Photo{}, photoID)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tidak ada foto yang ditemukan dengan ID yang diberikan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Foto berhasil dihapus"})
}

// getUserIDFromContext digunakan untuk mendapatkan ID pengguna dari konteks Gin.
func getUserIDFromContext(c *gin.Context) (uint, bool) {
	user, exists := c.Get("userID")
	if !exists {
		return 0, false
	}

	userID, ok := user.(uint)
	if !ok {
		return 0, false
	}
	return userID, true
}

