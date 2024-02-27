package databases

import (
	"os"
)

type dbconfig struct {
	Driver   string
	Host     string
	Port     string
	User     string
	Password string
	Db_name  string
	Sslmode  string
}

// DB_DRIVER = postgres
// DB_HOST = localhost
// DB_PORT = 5433
// DB_USER = postgres
// DB_PASSWORD = password
// DB_Name = blogreal
func config() dbconfig {
	db := dbconfig{
		Driver:   os.Getenv("DB_DRIVER"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Db_name:  os.Getenv("DB_NAME"),
		Sslmode:  os.Getenv("DB_SSLMODE"),
	}
	return db

}
