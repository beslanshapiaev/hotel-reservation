package middleware

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JWTQuthentication(c *fiber.Ctx) error {
	fmt.Println("--- jwt auth")

	token, ok := c.GetReqHeaders()["X-Api-Token"]
	if !ok {
		return fmt.Errorf("unauthorized not header")
	}
	if err := parseToken(token[0]); err != nil {
		return err
	}

	fmt.Println("token:", token)

	return nil
}

func parseToken(tokenStr string) error {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("invalid signing method", token.Header["alg"])
			return nil, fmt.Errorf("unauthorized not ok")
		}
		secret := os.Getenv("JWT_SECRET")
		fmt.Println("NEVER PRINT SECRET", secret)
		return []byte(secret), nil
	})
	if err != nil {
		fmt.Println("failed to parse token: ", err)
		return fmt.Errorf("unauthorized 123")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims)
	}
	return fmt.Errorf("unauthorized e")
}
