package models

import (
	"final-task/database"
	"final-task/app"
)

// createUser digunakan untuk membuat data pengguna baru dalam database.
func createUser(user *app.User) error {
	result := database.GetDB().Create(user)
	return result.Error
}

// FindUserByID digunakan untuk mencari data pengguna berdasarkan ID-nya dalam database.
func findUserByID(userID uint) (app.User, error) {
	var user app.User
	result := database.GetDB().First(&user, userID)
	return user, result.Error
}

// deleteUser digunakan untuk menghapus data pengguna dari database berdasarkan ID-nya.
func deleteUser(userID uint) error {
	result := database.GetDB().Delete(&app.User{}, userID)
	return result.Error
}

// updateUser digunakan untuk memperbarui data pengguna yang sudah ada dalam database.
func updateUser(user *app.User) error {
	result := database.GetDB().Save(user)
	return result.Error
}
