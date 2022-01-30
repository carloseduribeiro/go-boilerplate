package controller

import (
	"context"
	"errors"
	"time"

	"github.com/felipeagger/go-boilerplate/internal/domain"
	"github.com/felipeagger/go-boilerplate/internal/repository"
	"github.com/felipeagger/go-boilerplate/pkg/utils"
)

//CreateUser ...
func CreateUser(ctx context.Context, payload domain.Signup) error {

	userRepository := repository.NewGORMUserRepository()

	layOut := "2006/01/02"
	date, _ := time.Parse(layOut, payload.BirthDate)
	birthDate := date

	newUser := domain.User{
		ID:        utils.GeneratedID(),
		Name:      payload.Name,
		Email:     payload.Email,
		Password:  payload.Password,
		BirthDate: birthDate,
	}

	return userRepository.Create(ctx, newUser)
}

//SignInUser ...
func SignInUser(ctx context.Context, payload domain.Login) (domain.LoginResponse, error) {

	userRepository := repository.NewGORMUserRepository()

	user := userRepository.GetByEmail(ctx, payload.Email)

	if user.Password == payload.Password {
		return domain.LoginResponse{
			Token:   "jwt",
			Message: "Success",
		}, nil
	}

	return domain.LoginResponse{
		Message: "Email or Password incorrect!",
	}, errors.New("Unauthorized")
}