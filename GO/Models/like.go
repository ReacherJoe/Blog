
package entities

import "time"

type Like struct {
    ID        uint           `gorm:"primaryKey" json:"id"`
    CreatedAt time.Time      `json:"created_at"`
    UpdatedAt time.Time      `json:"updated_at"`
    UserID    uint           `json:"user_id"`
	PostID	  uint 			 `json:"post_id"`
}