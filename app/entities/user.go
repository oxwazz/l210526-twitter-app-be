package entities

import (
	"database/sql/driver"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/oxwazz/l210526-twitter-app-be/helpers"
	"golang.org/x/crypto/bcrypt"
	"reflect"
	"time"
)

type User struct {
	ID          helpers.NullString `gorm:"default:gen_random_uuid()" json:"id"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
	Name        string             `json:"name" validate:"required"`
	DateOfBirth helpers.CustomTime `json:"date_of_birth" validate:"required"`
	Email       helpers.NullString `json:"email" validate:"required_without_all=Username Phone,omitempty,email"`
	Username    helpers.NullString `json:"username" validate:"required_without_all=Email Phone"`
	Phone       helpers.NullString `json:"phone" validate:"required_without_all=Email Username,omitempty,e164"`
	Password    string             `json:"password" validate:"required"`
}

type UserForLogin struct {
	Name     string             `json:"name"`
	Email    helpers.NullString `json:"email" validate:"required_without_all=Username Phone,omitempty,email"`
	Username helpers.NullString `json:"username" validate:"required_without_all=Email Phone"`
	Phone    helpers.NullString `json:"phone" validate:"required_without_all=Email Username,omitempty,e164"`
	Password string             `json:"password" validate:"required"`
}

func (UserForLogin) TableName() string {
	return "users"
}

func ValidateValuer(field reflect.Value) interface{} {
	if valuer, ok := field.Interface().(driver.Valuer); ok {
		val, err := valuer.Value()
		if err == nil {
			if val == nil {
				return ""
			} else {
				return val
			}
		}
		// handle the error how you want
	}
	return nil
}

func (user User) IsValid() (bool, error) {
	v := validator.New()
	v.RegisterCustomTypeFunc(ValidateValuer, helpers.NullString{}, helpers.CustomTime{})
	err := v.Struct(user)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (user UserForLogin) IsValidForLogin() (bool, error) {
	v := validator.New()
	v.RegisterCustomTypeFunc(ValidateValuer, helpers.NullString{}, helpers.CustomTime{})
	err := v.Struct(user)
	if err != nil {
		return false, err
	}
	return true, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (user UserForLogin) ValidatePassword(password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return false, err
	}
	return true, nil
}

func NewUser(user User) (User, error) {
	_, err := user.IsValid()
	if err != nil {
		return user, err
	}

	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return user, err
	}

	u := User{
		ID:          user.ID,
		Name:        user.Name,
		DateOfBirth: user.DateOfBirth,
		Email:       user.Email,
		Username:    user.Username,
		Phone:       user.Phone,
		Password:    hashedPassword,
	}

	return u, nil
}

func (user UserForLogin) CreateJWT() (string, error) {
	j := jwt.New(jwt.SigningMethodHS256)
	claims := j.Claims.(jwt.MapClaims)
	claims["name"] = user.Name
	claims["username"] = user.Username.String
	claims["level"] = "application"
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	tokenJWT, err := j.SignedString([]byte("secret-key"))
	return tokenJWT, err
}
