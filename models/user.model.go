package models

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/oxwazz/l210526-twitter-app-be/helpers"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/oxwazz/l210526-twitter-app-be/db"

	"github.com/go-playground/validator/v10"
)

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

// github START
// https://stackoverflow.com/questions/69342582/how-do-i-bind-a-date-string-to-a-struct
// type CustomTime struct {
// 	time.Time
// }
type CustomTime struct {
	time.Time
}

func (t CustomTime) MarshalJSON() ([]byte, error) {
	date := t.Time.Format("2006-01-02")
	date = fmt.Sprintf(`"%s"`, date)
	return []byte(date), nil
}

func (t *CustomTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")

	date, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	t.Time = date
	return
}

// github END

// type Timestamp time.Time

// func (t *Timestamp) UnmarshalParam(src string) error {
// 	ts, err := time.Parse(time.RFC3339, src)
// 	*t = Timestamp(ts)
// 	return err
// }

// NullString is an alias for sql.NullString data type
type NullString struct {
	sql.NullString
}

// MarshalJSON for NullString
func (ns *NullString) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ns.String)
}

//UnmarshalJSON for NullString
func (ns *NullString) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &ns.String)
	ns.Valid = (err == nil)
	return err
}

type User struct {
	Id          string     `json:"id"`
	Name        string     `json:"name" validate:"required"`
	DateOfBirth CustomTime `json:"date_of_birth" validate:"required"`
	Email       NullString `json:"email" validate:"required_without_all=Username Phone,omitempty,email"`
	Username    NullString `json:"username" validate:"required_without_all=Email Phone"`
	Phone       NullString `json:"phone" validate:"required_without_all=Email Username,omitempty,e164"`
	Password    string     `json:"password" validate:"required"`
}

type ResponseUser struct {
	Id          int        `json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	Name        string     `json:"name" validate:"required"`
	DateOfBirth time.Time  `json:"date_of_birth" validate:"required"`
	Email       NullString `json:"email" validate:"required"`
	Username    NullString `json:"username" validate:"required"`
	Phone       NullString `json:"phone" validate:"required"`
	Password    string     `json:"password" validate:"required"`
}

func FetchAllUser() (Response, error) {
	var user ResponseUser
	var listUser []ResponseUser
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM users"

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(
			&user.Id,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.Name,
			&user.DateOfBirth,
			&user.Email,
			&user.Username,
			&user.Phone,
			&user.Password,
		)
		if err != nil {
			return res, err
		}

		listUser = append(listUser, user)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = listUser

	return res, nil
}

func Register(
	name string,
	dateOfBirth CustomTime,
	email NullString,
	username NullString,
	phone NullString,
	password string,
) (Response, error) {

	var res Response

	v := validator.New()
	v.RegisterCustomTypeFunc(ValidateValuer, NullString{})

	hashedPassword, err := helpers.HashPassword(password)
	if err != nil {
		return res, err
	}

	user := User{
		Name:        name,
		DateOfBirth: dateOfBirth,
		Email:       email,
		Username:    username,
		Phone:       phone,
		Password:    hashedPassword,
	}

	err = v.Struct(user)
	if err != nil {
		return res, err
	}

	con := db.CreateCon()

	sqlStatement := `
	INSERT INTO users (name, date_of_birth, email, username, phone, password)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING id`

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	var studentId string

	err = stmt.QueryRow(
		name,
		dateOfBirth.Format("2006-01-02"),
		email,
		username,
		phone,
		hashedPassword,
	).Scan(&studentId)
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]string{
		"last_inserted_id": studentId,
	}

	return res, nil
}

func UpdateUser(
	id int,
	name string,
	dateOfBirth CustomTime,
	email string,
	username string,
	phone string,
	password string,
) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := `
	UPDATE users
	SET name = $1, date_of_birth = $2, email = $3, username = $4, phone = $5, password = $6
	WHERE id = $7
	RETURNING id`

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	var updatedId string
	err = stmt.QueryRow(
		name,
		dateOfBirth.Format("2006-01-02"),
		email,
		username,
		phone,
		password,
		id,
	).Scan(&updatedId)
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]string{
		"rows_affected": updatedId,
	}

	return res, nil
}

func DeleteUser(id int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := `
	DELETE FROM users
	WHERE id = $1
	RETURNING id`

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	var deletedId string
	err = stmt.QueryRow(id).Scan(&deletedId)
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]string{
		"rows_affected": deletedId,
	}

	return res, nil
}
