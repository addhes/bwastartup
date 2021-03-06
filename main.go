package main

import (
	"bwastartup/auth"
	"bwastartup/handler"
	"bwastartup/user"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// nil = null alias gk ada isinya
	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	authService := auth.NewService()

	token, err := authService.ValidateToken("yJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxfQ.pR5mUyz1tm_Ni6-mCi-ankpmIwVifpJ0k_tNjbyp6p8")
	if err != nil {
		fmt.Println("error")
	}

	if token.Valid {
		fmt.Println("valid")
	} else {
		fmt.Println("INVALID")
	}

	userHandler := handler.NewUserHandler(userService, authService)

	//  membuat route /api/v1
	router := gin.Default()
	api := router.Group("/api/v1")

	//Membuat route post users
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailKetersedian)
	api.POST("/avatars", userHandler.UploadAvatar)
	// api.GET("/users/fetch", authMiddleware(authService, userService), userHandler.FetchUser)
	// menjalankan route
	router.Run()

}
