package helpers

import "golang.org/x/crypto/bcrypt"

// hashPassword menghasilkan hash dari password yang diberikan.
func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// checkPasswordHash memeriksa apakah password sesuai dengan hash yang diberikan.
func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
