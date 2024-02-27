
package entities

import "time"

type Reply struct {
    ID        uint           `gorm:"primaryKey" json:"id"`
	Text      string		 `json:"text"`
    CreatedAt time.Time      `json:"created_at"`
    UpdatedAt time.Time      `json:"updated_at"`
    UserID    uint           `json:"user_id"`
	CommentID uint 			 `json:"comment_id"`
}