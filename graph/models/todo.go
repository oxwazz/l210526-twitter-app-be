package models

import "github.com/oxwazz/twitter/helpers"

type Todo struct {
	ID     string             `json:"id"`
	Text   string             `json:"text"`
	Done   bool               `json:"done"`
	UserID helpers.NullString `json:"user_id"`
}
