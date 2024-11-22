package helpers

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"

	"go-rest-api-jwt/models"
)

var mySigningKey = []byte("mysecretkey")

type MyCustomClaims struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

// CreateToken is a function that takes a pointer to models.User and returns a string and an error.
// The string is a JWT that contains the user's ID, name, and email, and is signed with a secret key.
// The error is returned if there is an error generating the JWT.
// The JWT is valid for 24 hours, and is not valid before the current time.
func CreateToken(user *models.User) (string, error) {
	claims := MyCustomClaims{
		user.ID,
		user.Name,
		user.Email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)

	return ss, err
}

// ValidateToken takes a JWT string as input and returns the claims if the token is valid,
// or an error if the token is invalid or unauthorized. It parses the token using the
// custom claims structure MyCustomClaims and verifies it using a predefined signing key.
func ValidateToken(tokenString string) (any, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("Unauthorized")
	}

	claims, ok := token.Claims.(*MyCustomClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("Unauthorized")
	}

	return claims, nil
}
