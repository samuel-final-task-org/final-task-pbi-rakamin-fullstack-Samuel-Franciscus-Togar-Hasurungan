package models

import (
	"final-task/database"
	"final-task/app"
)


// createPhoto digunakan untuk membuat data Photo baru dalam database.
func createPhoto(photo *app.Photo) error {
	result := database.GetDB().Create(photo)
	return result.Error
}

// findPhotoByID digunakan untuk mencari data Photo berdasarkan ID-nya dalam database.
func findPhotoByID(photoID uint) (app.Photo, error) {
	var photo app.Photo
	result := database.GetDB().First(&photo, photoID)
	return photo, result.Error
}

// updatePhoto digunakan untuk memperbarui data Photo yang sudah ada dalam database.
func updatePhoto(photo *app.Photo) error {
	result := database.GetDB().Save(photo)
	return result.Error
}

// deletePhoto digunakan untuk menghapus data Photo dari database berdasarkan ID-nya.
func deletePhoto(photoID uint) error {
	result := database.GetDB().Delete(&app.Photo{}, photoID)
	return result.Error
}
