
package entities

import "time"

type Post_Categorie struct {
    ID        uint           `gorm:"primaryKey" json:"id"`
    CreatedAt    time.Time       `json:"created_at"`
    UpdatedAt    time.Time       `json:"updated_at"`
	PostID	     uint 			 `json:"post_id"`
	CategorieID  uint		     `json:"categorie_id"`
}