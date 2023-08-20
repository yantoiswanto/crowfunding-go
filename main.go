package main

import (
	"crowfunding/handler"
	repository "crowfunding/repositories"
	"crowfunding/services"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:password@tcp(127.0.0.1:3306)/db_crowfunding?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	authService := services.NewService()
	log, err := authService.GenerateToken(1)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(log)
	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()
	api := router.Group("api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvaibality)
	api.POST("/avatars", userHandler.UploadAvatar)

	router.Run()

}
