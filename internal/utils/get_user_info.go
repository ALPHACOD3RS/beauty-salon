package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type UserInfo struct {
    UserID string
    Email  string
    Role   string
}



func GetUserInfoFromJWT(c *fiber.Ctx) UserInfo {
    userToken := c.Locals("user").(*jwt.Token)
    claims := userToken.Claims.(jwt.MapClaims)
    
    return UserInfo{
        UserID: claims["user_id"].(string),
        Email:  claims["email"].(string),
        Role:   claims["role"].(string),
    }
}