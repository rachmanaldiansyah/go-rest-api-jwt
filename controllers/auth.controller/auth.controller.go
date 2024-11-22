package authcontroller

import (
	"encoding/json"
	"net/http"

	"go-rest-api-jwt/config"
	"go-rest-api-jwt/helpers"
	"go-rest-api-jwt/models"
)

// Register is a function that handles the registration of a new user.
// It takes two parameters, http.ResponseWriter and http.Request.
// It returns a JSON response with a status of 200 if the registration is successful,
// or a status of 400 if the password and password confirmation do not match,
// or a status of 500 if there is an error with the database.
func Register(w http.ResponseWriter, r *http.Request) {
	var register models.Register

	if err := json.NewDecoder(r.Body).Decode(&register); err != nil {
		helpers.Response(w, 400, err.Error(), nil)
		return
	}

	defer r.Body.Close()

	if register.Password != register.PasswordConfirm {
		helpers.Response(w, 400, "Password not match", nil)
		return
	}

	passwordHash, err := helpers.HashPassword(register.Password)
	if err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	user := models.User{
		Name:     register.Name,
		Email:    register.Email,
		Password: passwordHash,
	}

	if err := config.DB.Create(&user).Error; err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	helpers.Response(w, 200, "Register successfully", nil)
}

// Login is a function that handles the login of an existing user.
// It takes two parameters, http.ResponseWriter and http.Request.
// It returns a JSON response with a status of 200 if the login is successful,
// or a status of 404 if the email or password is incorrect,
// or a status of 500 if there is an error with the database.
//
// The response body will contain the JSON Web Token (JWT) of the user.
func Login(w http.ResponseWriter, r *http.Request) {
	var login models.Login

	if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	var user models.User
	if err := config.DB.First(&user, "email = ?", login.Email).Error; err != nil {
		helpers.Response(w, 404, "Wrong email or password", nil)
		return
	}

	if err := helpers.VerifyPassword(user.Password, login.Password); err != nil {
		helpers.Response(w, 404, "Wrong email or password", nil)
		return
	}

	token, err := helpers.CreateToken(&user)
	if err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	helpers.Response(w, 200, "Login Successfully", token)
}
