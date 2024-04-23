package api

import (
	"errors"
	"fmt"

	"github.com/beslanshapiaev/hotel-reservation/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	userStore db.UserStore
}

func NewAuthHandler(userStore db.UserStore) *AuthHandler {
	return &AuthHandler{
		userStore: userStore,
	}
}

type AuthParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *AuthHandler) HandleAuthenticate(c *fiber.Ctx) error {
	var authParams AuthParams

	if err := c.BodyParser(&authParams); err != nil {
		return err
	}

	user, err := h.userStore.GetUserByEmail(c.Context(), authParams.Email)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return fmt.Errorf("user not found")
		}
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.EncryptedPassword), []byte(authParams.Password)); err != nil {
		return fmt.Errorf("invalid creadetials")
	}
	fmt.Println("user authenticated")

	return nil
}
