package repositories

import (
	"fmt"
	"l210526-twitter-app-be/app/entities"
	"l210526-twitter-app-be/app/entities/databases"
	"l210526-twitter-app-be/helpers"
)

func FetchAllUser[U []*entities.User | []entities.User]() (U, error) {
	db := databases.CreateConnection()
	var users U
	result := db.Order("created_at desc").Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func FetchUserByID[T *entities.User | entities.User](ID string) (T, error) {
	db := databases.CreateConnection()
	var user T
	result := db.Where("id = ?", ID).First(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

//func FetchAllUser() ([]entities.User, error) {
//	db := databases.CreateConnection()
//	var users []entities.User
//	result := db.Order("created_at desc").Find(&users)
//	if result.Error != nil {
//		return users, result.Error
//	}
//	return users, nil
//}

func Register(user entities.User) (entities.User, error) {
	db := databases.CreateConnection()

	newUser, err := entities.NewUser(user)
	if err != nil {
		return user, err
	}

	fmt.Println(333300, newUser)
	result := db.Create(&newUser) // pass pointer of data to Create
	if result.Error != nil {
		return entities.User{}, result.Error
	}

	return newUser, nil
}

func Login(email, username, phone helpers.NullString, password string) (string, error) {
	db := databases.CreateConnection()

	user := entities.UserForLogin{
		Email:    email,
		Username: username,
		Phone:    phone,
		Password: password,
	}

	_, err := user.IsValidForLogin()
	if err != nil {
		return "", err
	}

	// fetch user by email, username, or phone
	result := db.Where(entities.UserForLogin{
		Email:    user.Email,
		Username: user.Username,
		Phone:    user.Phone,
	}).First(&user)
	// if not found handle error
	if result.Error != nil {
		return "", result.Error
	}

	// check bcrypt password
	_, err = user.ValidatePassword(password)
	// if not valid handle error
	if err != nil {
		return "", err
	}

	// send token jwt
	token, err := user.CreateJWT()
	if err != nil {
		return "", err
	}
	return token, nil
}
