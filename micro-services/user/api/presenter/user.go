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

type UserRefreshRequest struct {
	RefreshToken string `json:"refreshToken" validate:"required"`
}

type UserRefreshResponse struct {
	AccessToken string `json:"accessToken"`
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



