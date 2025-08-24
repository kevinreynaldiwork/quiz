package controllers

import (
	"net/http"
	"time"

	"Quiz/database"
	"Quiz/repository"
	"Quiz/structs"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("SECRET_KEY")

// Login user dari DB
func Login(c *gin.Context) {
	var input structs.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Ambil user dari DB via repository
	dbUser, err := repository.GetUserByUsername(database.DB, input.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User tidak ditemukan"})
		return
	}

	// Validasi password pakai bcrypt
	if bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(input.Password)) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Password salah"})
		return
	}

	// Generate token JWT
	expiration := time.Now().Add(2 * time.Hour)
	claims := &jwt.RegisteredClaims{
		Subject:   dbUser.Username,
		ExpiresAt: jwt.NewNumericDate(expiration),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, _ := token.SignedString(jwtKey)

	c.JSON(http.StatusOK, gin.H{
		"token": tokenStr,
		"user": gin.H{
			"id":         dbUser.ID,
			"username":   dbUser.Username,
			"created_by": dbUser.CreatedBy,
		},
	})
}

// (Optional) Register user baru
func Register(c *gin.Context) {
	var input structs.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Hash password
	hashed, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	input.Password = string(hashed)

	// Insert user ke DB
	err := repository.InsertUser(database.DB, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}
