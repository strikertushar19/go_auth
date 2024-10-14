package main

import (
	"go_auth/controllers"
	"go_auth/initializers"
	"go_auth/middleware"
	"go_auth/docs"  // Swagger docs should be imported like this

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	// Load environment variables and connect to the database
	initializers.LoadEnv()
	initializers.ConnectDB()
}

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name    API Support
// @contact.url     http://www.swagger.io/support
// @contact.email   support@swagger.io

// @license.name    Apache 2.0
// @license.url     http://www.apache.org/licenses/LICENSE-2.0.html

// @host            localhost:8080
// @BasePath        /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description   OpenAPI
// @externalDocs.url           https://swagger.io/resources/open-api/
func main() {
	// Use the docs package
	docs.SwaggerInfo.Title = "Swagger Example API"

	router := gin.Default()

	// Authentication routes
	router.POST("/auth/signup", controllers.CreateUser)
	router.POST("/auth/login", controllers.Login)
	router.GET("/user/profile", middlewares.CheckAuth, controllers.GetUserProfile)

	// Swagger route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Health check route
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	// Start the server
	router.Run(":8080")
}
