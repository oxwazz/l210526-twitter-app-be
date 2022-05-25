package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"l210526-twitter-app-be/db"

	"github.com/go-playground/validator/v10"
)

type PropertyMap map[string]interface{}

func (p PropertyMap) Value() (driver.Value, error) {
	j, err := json.Marshal(p)
	return j, err
}

func (p *PropertyMap) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return errors.New("Type assertion .([]byte) failed.")
	}

	var i interface{}
	err := json.Unmarshal(source, &i)
	if err != nil {
		return err
	}

	*p, ok = i.(map[string]interface{})
	if !ok {
		return errors.New("Type assertion .(map[string]interface{}) failed.")
	}

	return nil
}

type Tweet struct {
	Id         string      `json:"id"`
	Content    string      `json:"content" validate:"required"`
	Attachment PropertyMap `json:"attachment" validate:"required"`
	UserId     string      `json:"user_id"`
}

type ResponseTweet struct {
	Id         int         `json:"id"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
	Content    string      `json:"content" validate:"required"`
	Attachment PropertyMap `json:"attachment" validate:"required"`
	UserId     int         `json:"user_id"`
}

func FetchAllTweet() (Response, error) {
	var tweet ResponseTweet
	var listTweet []ResponseTweet
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM tweets"

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(
			&tweet.Id,
			&tweet.CreatedAt,
			&tweet.UpdatedAt,
			&tweet.Content,
			&tweet.Attachment,
			&tweet.UserId,
		)
		if err != nil {
			return res, err
		}

		listTweet = append(listTweet, tweet)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = listTweet

	return res, nil
}

func StoreTweet(
	userId int,
	content string,
	attachment PropertyMap,
) (Response, error) {

	var res Response

	v := validator.New()

	tweet := Tweet{
		Content:    content,
		Attachment: attachment,
	}

	err := v.Struct(tweet)
	if err != nil {
		return res, err
	}

	con := db.CreateCon()

	sqlStatement := `
	INSERT INTO tweets (content, attachment, user_id)
	VALUES ($1, $2, $3)
	RETURNING id`

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	var tweetId string

	err = stmt.QueryRow(
		content,
		attachment,
		userId,
	).Scan(&tweetId)
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]string{
		"last_inserted_id": tweetId,
	}

	return res, nil
}

func UpdateTweet(
	id int,
	content string,
	attachment PropertyMap,
) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := `
	UPDATE tweets
	SET content = $1, attachment = $2
	WHERE id = $3
	RETURNING id`

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	var updatedId string
	err = stmt.QueryRow(
		content,
		attachment,
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

func DeleteTweet(id int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := `
	DELETE FROM tweets
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
