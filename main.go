package main

import (
	//"example/restapi/album"
	database "api/db"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	loadDatabase()
	// serverApplication()
}

func loadDatabase() {
	database.Connect()
	// database.Database.AutoMigrate(&model.User{})
	// database.Database.AutoMigrate(&model.Entry{})
}
func loadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// func serverApplication() {
// 	router := gin.Default()

// 	publicRoutes := router.Group("/auth")
// 	publicRoutes.POST("/register", controller.Register)
// 	publicRoutes.POST("/login", controller.Login)
// 	protectedRoutes := router.Group("/api")
// 	protectedRoutes.Use(middleware.JWTAuthMiddleware())
// 	protectedRoutes.POST("/entry", controller.AddEntry)
// 	protectedRoutes.GET("/entry", controller.GetAllEntries)
// 	protectedRoutes.GET("/entryy", controller.GetAllEntries)

// 	router.Run(":8000")

// 	fmt.Println("server running on port 8000")

// }

// func main(){
// 	Generic()
// 	GenericCall()
// }
