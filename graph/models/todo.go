package models

import "l210526-twitter-app-be/helpers"

type Todo struct {
	ID     string             `json:"id"`
	Text   string             `json:"text"`
	Done   bool               `json:"done"`
	UserID helpers.NullString `json:"user_id"`
}
