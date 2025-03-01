package handlers

import (
	"Phinance/database"
	"Phinance/dto"
	"Phinance/models"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Login(c *gin.Context) {
	var user models.User
	var userDTO dto.UserLoginDTO
	var token string

	if err := c.ShouldBindJSON(&userDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.First(&user, "name = ?", userDTO.Name)

	if userDTO.Password != user.Password {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid username or password"})
		return
	}

	user = models.User{
		Name:     userDTO.Name,
		Password: userDTO.Password,
	}

	token = createToken(user)
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func createToken(user models.User) string {
	// Create the Claims
	claims := jwt.MapClaims{
		"name": user.Name,
		"exp":  time.Now().Add(time.Second * 240).Unix(),
	}
	// generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		fmt.Println(err)
	}
	return t
}
