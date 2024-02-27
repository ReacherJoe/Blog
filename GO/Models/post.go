
package entities

import "time"

type Post struct {
    ID        uint           `gorm:"primaryKey" json:"id"`
    Title     string         `json:"title"`
    Body      string         `json:"body"`
    Photo     string         `json:"photo"`
    CreatedAt time.Time      `json:"created_at"`
    UpdatedAt time.Time      `json:"updated_at"`
    UserID    uint           `json:"user_id"`
	Likes     []Like	     `gorm:"foreignKey:PostID" json:"likes"`
	Comments  []Comment	     `gorm:"foreignKey:PostID" json:"comments"`
    Post_Categories []Post_Categorie    `gorm:"foreignKey:PostID" json:"post_categories"`
}