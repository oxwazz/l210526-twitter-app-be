package models

import (
	"net/http"
	"time"

	"github.com/oxwazz/twitter/db"

	"github.com/go-playground/validator/v10"
)

// github START
// https://stackoverflow.com/questions/69342582/how-do-i-bind-a-date-string-to-a-struct
// type CustomTime struct {
// 	time.Time
// }
// type CustomTime struct {
// 	time.Time
// }

// func (t CustomTime) MarshalJSON() ([]byte, error) {
// 	date := t.Time.Format("2006-01-02")
// 	date = fmt.Sprintf(`"%s"`, date)
// 	return []byte(date), nil
// }

// func (t *CustomTime) UnmarshalJSON(b []byte) (err error) {
// 	s := strings.Trim(string(b), "\"")

// 	date, err := time.Parse("2006-01-02", s)
// 	if err != nil {
// 		return err
// 	}
// 	t.Time = date
// 	return
// }

// github END

// type Timestamp time.Time

// func (t *Timestamp) UnmarshalParam(src string) error {
// 	ts, err := time.Parse(time.RFC3339, src)
// 	*t = Timestamp(ts)
// 	return err
// }

type Follower struct {
	Id         string `json:"id"`
	FollowedId string `json:"followed_id" validate:"required"`
	FollowerId string `json:"follower_id" validate:"required"`
}

type ResponseFollower struct {
	Id         int       `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	FollowedId int       `json:"followed_id" validate:"required"`
	FollowerId int       `json:"follower_id" validate:"required"`
}

func FetchAllFollower() (Response, error) {
	var follower ResponseFollower
	var listFollower []ResponseFollower
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM followers"

	rows, err := con.Query(sqlStatement)
	defer rows.Close()
	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(
			&follower.Id,
			&follower.UpdatedAt,
			&follower.CreatedAt,
			&follower.FollowedId,
			&follower.FollowerId,
		)
		if err != nil {
			return res, err
		}

		listFollower = append(listFollower, follower)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = listFollower

	return res, nil
}

func StoreFollower(
	followedId string,
	followerId string,
) (Response, error) {

	var res Response

	v := validator.New()

	follower := Follower{
		FollowedId: followedId,
		FollowerId: followerId,
	}

	err := v.Struct(follower)
	if err != nil {
		return res, err
	}

	con := db.CreateCon()

	sqlStatement := `
	INSERT INTO followers (followed_id, follower_id)
	VALUES ($1, $2)
	RETURNING id`

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	var id string

	err = stmt.QueryRow(
		followedId,
		followerId,
	).Scan(&id)
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]string{
		"last_inserted_id": id,
	}

	return res, nil
}

func UpdateFollower(
	id int,
	followedId string,
	followerId string,
) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := `
	UPDATE followers
	SET followed_id = $1, follower_id = $2
	WHERE id = $3
	RETURNING id`

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	var updatedId string
	err = stmt.QueryRow(
		followedId,
		followerId,
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

func DeleteFollower(id int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := `
	DELETE FROM followers
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
