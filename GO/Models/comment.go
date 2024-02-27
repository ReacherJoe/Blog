
package entities

import "time"

type Comment struct {
    ID        uint           `gorm:"primaryKey" json:"id"`
	Text      string		 `json:"text"`
    CreatedAt time.Time      `json:"created_at"`
    UpdatedAt time.Time      `json:"updated_at"`
    UserID    uint           `json:"user_id"`
	PostID	  uint 			 `json:"post_id"`
	Replys    []Reply        `gorm:"foreignKey:CommentID" json:"replys"`
	
}