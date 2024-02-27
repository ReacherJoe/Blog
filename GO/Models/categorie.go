
package entities

import "time"

type Categorie struct {
    ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string		 `json:"name"`
    CreatedAt time.Time      `json:"created_at"`
    UpdatedAt time.Time      `json:"updated_at"`
	Post_Categories []Post_Categorie    `gorm:"foreignKey:CategorieID" json:"post_categories"`
}