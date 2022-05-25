package usecases

import (
	"fmt"
	"github.com/oxwazz/twitter/app/entities"
	"github.com/oxwazz/twitter/app/entities/repositories"
	"github.com/oxwazz/twitter/helpers"
)

func FetchAllUser() ([]entities.User, error) {
	users, err := repositories.FetchAllUser[[]entities.User]()
	if err != nil {
		return []entities.User{}, err
	}
	return users, nil
}

func Register(user entities.User) (entities.User, error) {
	user, err := repositories.Register(user)
	if err != nil {
		return entities.User{}, err
	}
	return user, nil
}

func Login(email, username, phone helpers.NullString, password string) (string, error) {
	if email.String == "" {
		fmt.Println(333)
	}

	token, err := repositories.Login(email, username, phone, password)
	if err != nil {
		return "", err
	}
	return token, nil
}
