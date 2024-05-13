package helpers

import (
	"fmt"
	"time"
    "github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("your_secret_key")

// generateToken digunakan untuk menghasilkan token JWT berdasarkan userID.
func generateToken(userID uint) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// Claims untuk menyimpan struktur data dari token JWT.
type Claims struct {
    UserID uint `json:"user_id"`
    jwt.StandardClaims
}

// verifyToken digunakan untuk memverifikasi validitas token JWT dan mengembalikan klaim-klaim yang ada di dalamnya.
func verifyToken(tokenString string) (*Claims, error) {
    claims := &Claims{}
    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })
    
    if err != nil {
        return nil, err
    }
    if !token.Valid {
        return nil, fmt.Errorf("invalid token")
    }
    return claims, nil
}
