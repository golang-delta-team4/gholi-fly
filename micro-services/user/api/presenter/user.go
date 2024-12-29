package presenter

import (
	"errors"
	"regexp"
)

type UserSignUpRequest struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required"`
	Password  string `json:"password" validate:"required"`
}

type UserSignInRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserSignInResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type GetUserResponse struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

type PaginationQuery struct {
    Page int `query:"page" default:"1" validate:"gt=0"`
    Size int `query:"size" default:"10" validate:"gt=0"`
}

func EmailValidation(email string) error {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	emailMatched, err := regexp.MatchString(emailRegex, email)
	if err != nil {
		return err
	}
	if !emailMatched {
		return errors.New("invalid email format")
	}
	return nil
}

func PasswordValidation(password string) bool {
	uppercaseRegex := regexp.MustCompile(`[A-Z]`)
	lowercaseRegex := regexp.MustCompile(`[a-z]`)
	digitRegex := regexp.MustCompile(`[0-9]`)
	specialCharRegex := regexp.MustCompile(`[#!?@$%^&*\\-]`)
	minLength := len(password) >= 8
	return uppercaseRegex.MatchString(password) &&
		lowercaseRegex.MatchString(password) &&
		digitRegex.MatchString(password) &&
		specialCharRegex.MatchString(password) &&
		minLength

}
