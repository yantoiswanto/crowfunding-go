package main

import (
	"crowfunding/handler"
	"crowfunding/helper"
	repository "crowfunding/repositories"
	"crowfunding/services"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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
	campaignRepository := repository.NewCampaignRepository(db)

	userService := services.NewUserService(userRepository)
	campaignService := services.NewCampaignService(campaignRepository)
	authService := services.NewService()

	userHandler := handler.NewUserHandler(userService, authService)
	campaignHander := handler.NewCampaignHandler(campaignService)

	router := gin.Default()
	router.Static("/images", "./images/")
	api := router.Group("api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvaibality)
	api.POST("/avatars", authMiddleware(authService, userService), userHandler.UploadAvatar)

	api.GET("/campaigns", campaignHander.GetCampaigns)

	router.Run()

}
func authMiddleware(authService services.AuthService, userService services.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorization", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorization", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorization", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := claim["user_id"].(float64)

		user, err := userService.GetUserByID(int(userID))
		if err != nil {
			response := helper.APIResponse("Unauthorization", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)

	}
}
