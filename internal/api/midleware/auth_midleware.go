package midleware

import (
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMidleware(c *fiber.Ctx) error{
	tokenString := c.Get("Authorization")


	if tokenString == "" {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "msg": "Missing or malformed JWT",
        })
    }

	parts := strings.Split(tokenString, " ")
    if len(parts) == 2 && parts[0] == "Bearer" {
        tokenString = parts[1]
    }

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fiber.NewError(fiber.StatusUnauthorized, "Unexpected signing method")
        }
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if err != nil{
		return c.Status(400).JSON(fiber.Map{
			"msg": "Invalid token",
		})
	}

	if !token.Valid {
		return c.Status(400).JSON(fiber.Map{
			"msg": "Invalid token",
		})
	}
	claims, ok := token.Claims.(jwt.MapClaims)
    if !ok || !token.Valid {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "msg": "Invalid token claims",
        })
    }

    exp, ok := claims["exp"].(float64)
    if !ok {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "msg": "Invalid token expiration",
        })
    }

    if exp < float64(time.Now().Unix()) {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "msg": "Token has expired",
        })
    }


	return c.Next()

}

