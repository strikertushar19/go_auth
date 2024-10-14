
// controllers/authController.go

package controllers

import (
	"go_auth/initializers"
	"go_auth/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// @Summary      Create User
// @Description  Registers a new user
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        user  body     models.AuthInput  true  "User Signup Info"
// @Success      200  {object}  models.User
// @Failure      400  {object}  map[string]string "error": "Bad Request"
// @Router       /auth/signup [post]
func CreateUser(c *gin.Context) {

	var authInput models.AuthInput

	if err := c.ShouldBindJSON(&authInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userFound models.User
	initializers.DB.Where("username = ?", authInput.Username).Find(&userFound)

	if userFound.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username already used"})
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(authInput.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Username: authInput.Username,
		Password: string(passwordHash),
	}

	initializers.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})

}
// @Summary      User Login
// @Description  Logs in an existing user and returns a JWT token
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        user  body     models.AuthInput  true  "User Login Info"
// @Success      200  {object}  map[string]string "token": "JWT Token"
// @Failure      400  {object}  map[string]string "error": "Bad Request"
// @Router       /auth/login [post]
func Login(c *gin.Context) {

	var authInput models.AuthInput

	if err := c.ShouldBindJSON(&authInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userFound models.User
	initializers.DB.Where("username=?", authInput.Username).Find(&userFound)

	if userFound.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userFound.Password), []byte(authInput.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid password"})
		return
	}

	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  userFound.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := generateToken.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to generate token"})
	}

	c.JSON(200, gin.H{
		"token": token,
	})
}
// @Summary      Get User Profile
// @Description  Returns the current user's profile
// @Tags         Auth
// @Produce      json
// @Success      200  {object}  map[string]interface{} "user": "Current User Data"
// @Router       /user/profile [get]
func GetUserProfile(c *gin.Context) {

	user, _ := c.Get("currentUser")

	c.JSON(200, gin.H{
		"user": user,
	})
}