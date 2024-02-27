package databases

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres" // Import PostgreSQL driver
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	if config().Driver == "postgres" { // Check if the driver is PostgreSQL
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
			config().Host, config().User, config().Password, config().Db_name, config().Port, config().Sslmode)

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

		if err != nil {
			log.Println("Error connecting to the database:", err)
			return nil, err
		}

		return db, nil
	}

	return nil, fmt.Errorf("unsupported database driver: %s", config().Driver)
}
