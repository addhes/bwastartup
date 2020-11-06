package main

import (
	"bwastartup/handler"
	"bwastartup/user"
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

	userHandler := handler.NewUserHandler(userService)

	//  membuat route /api/v1
	router := gin.Default()
	api := router.Group("/api/v1")

	//Membuat route post users
	api.POST("/users", userHandler.RegisterUser)
	// menjalankan route
	router.Run()

	// fmt.Println("Connetion to database success")

	// var users []user.User //variable users adalah struct dari User(user/entity.go/User)
	// db.Find(&users)       // mencari variable users

	// for _, user := range users {
	// 	fmt.Println(user.Name)
	// 	fmt.Println(user.Email)
	// 	fmt.Println("=======")
	// }

}
