package types

import (
	"fmt"
	"regexp"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

const (
	bcryptCost             = 12
	minimumFirstNameLength = 2
	minimumLastNameLength  = 2
	minimumPasswordLength  = 7
)

type UpdateUserParams struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func (p UpdateUserParams) ToBSON() bson.M {
	result := bson.M{}
	if len(p.FirstName) > 0 {
		result["firstName"] = p.FirstName
	}
	if len(p.LastName) > 0 {
		result["lastName"] = p.LastName
	}
	return result
}

type CreateUserParams struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (params CreateUserParams) Validate() map[string]string {
	errors := map[string]string{}
	if len(params.FirstName) < minimumFirstNameLength {
		errors["firstName"] = fmt.Sprintf("firstName length should be at least %d characters", minimumFirstNameLength)
	}
	if len(params.LastName) < minimumLastNameLength {
		errors["lastName"] = fmt.Sprintf("lastName length should be at least %d characters", minimumLastNameLength)
	}
	if len(params.Password) < minimumPasswordLength {
		errors["password"] = fmt.Sprintf("password length should be at least %d characters", minimumPasswordLength)
	}
	if !isEmailValid(params.Email) {
		errors["email"] = fmt.Sprintf("email is invalid")
	}
	return errors
}

func isEmailValid(e string) bool {
	var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if len(e) < 3 && len(e) > 254 {
		return false
	}
	return emailRegex.MatchString(e)
}

type User struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName         string             `bson:"firstName" json:"firstName"`
	LastName          string             `bson:"lastName" json:"lastName"`
	Email             string             `bson:"email" json:"email"`
	EncryptedPassword string             `bson:"encryptedPassword" json:"-"`
}

func NewUserFromParams(params CreateUserParams) (*User, error) {
	encpw, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcryptCost)
	if err != nil {
		return nil, err
	}
	return &User{
		FirstName:         params.FirstName,
		LastName:          params.LastName,
		Email:             params.Email,
		EncryptedPassword: string(encpw),
	}, nil
}
