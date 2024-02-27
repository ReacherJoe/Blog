package migrations

import (
	"fmt"
	databases "gorm_api/Databases"
	models "gorm_api/Models"
	"log"
)

func Migrate(action string, model_name string) {
	db, err := databases.Connect()
	if err != nil {
		log.Println(err)
	}

	if action == "migrate" {
		if model_name == "all" {
			db.AutoMigrate(&models.User{})
			db.AutoMigrate(&models.Post{})
			db.AutoMigrate(&models.Like{})
			db.AutoMigrate(&models.Comment{})
			db.AutoMigrate(&models.Reply{})
			db.AutoMigrate(&models.Categorie{})
			db.AutoMigrate(&models.Post_Categorie{})
			fmt.Println("All models migrate success...")
		} else if model_name == "user" {
			db.AutoMigrate(&models.User{})
			fmt.Println("User model migrate success...")

		} else if model_name == "post"{
			db.AutoMigrate(&models.Post{})
			fmt.Println("Post model migrate success...")

		}else if model_name == "like"{
			db.AutoMigrate(&models.Like{})
			fmt.Println("Like model migrate success...")

		}else if model_name == "comment"{
			db.AutoMigrate(&models.Comment{})
			fmt.Println("Comment model migrate success...")

		}else if model_name == "reply"{
			db.AutoMigrate(&models.Reply{})
			fmt.Println("Reply model migrate success...")

		}else if model_name == "categories"{
			db.AutoMigrate(&models.Categorie{})
			fmt.Println("Categories model migrate success...")

		}else if model_name == "post_categories"{
			db.AutoMigrate(&models.Post_Categorie{})
			fmt.Println("Post_Categories model migrate success...")

		}else{
			fmt.Printf("No match model name with %s\n", model_name)
		}
	} else if action == "drop" {
		if model_name == "all" {
			db.Migrator().DropTable(&models.User{})
			db.Migrator().DropTable(&models.Post{})
			db.Migrator().DropTable(&models.Like{})
			db.Migrator().DropTable(&models.Comment{})
			db.Migrator().DropTable(&models.Reply{})
			db.Migrator().DropTable(&models.Categorie{})
			db.Migrator().DropTable(&models.Post_Categorie{})

			fmt.Println("All models drop success...")
		} else if model_name == "user" {
			db.Migrator().DropTable(&models.User{})
			fmt.Println("User model drop success...")

		} else if model_name == "post" {
			db.Migrator().DropTable(&models.Post{})
			fmt.Println("Post model drop success...")

		} else if model_name == "like" {
			db.Migrator().DropTable(&models.Like{})
			fmt.Println("Like model drop success...")

		} else if model_name == "comment" {
			db.Migrator().DropTable(&models.Post{})
			fmt.Println("Post model drop success...")

		} else if model_name == "reply" {
			db.Migrator().DropTable(&models.Reply{})
			fmt.Println("Reply model drop success...")

		} else if model_name == "categories" {
			db.Migrator().DropTable(&models.Categorie{})
			fmt.Println("Categories model drop success...")

		} else if model_name == "post_categories" {
			db.Migrator().DropTable(&models.Post_Categorie{})
			fmt.Println("Post_Categories model drop success...")

		} else {
			fmt.Printf("No match model name with %s\n", model_name)
		}
	} else {
		fmt.Println("Unknown action")
	}

}
