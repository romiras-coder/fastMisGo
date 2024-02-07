package main

import (
	"api/api_router"
	configuration "api/config"
	"api/database"
	model "api/models"
	"fmt"
	"net/http"

	_ "api/docs"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var cfg = configuration.ReadConfig()

func main() {
	loadDatabase()
	serverApplication()
}

func loadDatabase() {
	database.Connect(cfg)
	database.Database.AutoMigrate(&model.User{})
	database.Database.AutoMigrate(&model.Entry{})
}

// @title           Test API in GO
// @version         1.0
// @description     Тестовая API-шка для поиграться с go postgres и прочими штуками
// @termsOfService  https://блаблабла

// @contact.name   Moskotlinov Roman
// @contact.url    https://github.com/romiras-coder
// @contact.email  test@test.ru

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8899
// @BasePath  /api/v1

func serverApplication() {
	router := gin.Default()
	router.Use(CORSMiddleware())

	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/api/v1/docs/index.html#/")
	})

	v1 := router.Group("/api/v1")
	{
		v1.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		authRoutes := v1.Group("/auth")
		authRoutes.POST("/register", api_router.Register)
		authRoutes.POST("/login", api_router.Login)
	}

	// publicRoutes := v1.Group("/auth")
	// publicRoutes.POST("/register", api_router.Register)
	// publicRoutes.POST("/login", api_router.Login)

	// protectedRoutes := router.Group("/auth")
	// protectedRoutes.Use(middleware.JWTAuthMiddleware())
	// protectedRoutes.POST("/entry", api_router.AddEntry)
	// protectedRoutes.GET("/entry", api_router.GetAllEntries)
	// protectedRoutes.GET("/entryy", api_router.GetAllEntries)
	router.Run(fmt.Sprintf(":%d", cfg.ApiService.Port))
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
