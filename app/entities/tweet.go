package entities

import (
	"time"
)

type Tweet struct {
	ID        string    `json:"id" gorm:"default:gen_random_uuid()"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Content   string    `json:"content" `
	//Attachment AttachmentWrapper `json:"attachment" `
	UserID string `json:"user_id"`
}

//type AttachmentWrapper struct {
//	Image string `json:"image"`
//}
