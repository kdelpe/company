package mysql

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var db *sql.DB

func Connection() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
	cfg := mysql.Config{
		User:      os.Getenv("DB_USERNAME"),
		Passwd:    os.Getenv("DB_PASSWORD"),
		Net:       "tcp",
		Addr:      "127.0.0.1:3306",
		DBName:    "company",
		ParseTime: true,
	}

	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatalf("Error connecting to database, %s", err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatalf("Lost Connection to the databse, %s", pingErr)
	}

	fmt.Println("Successful connected to MySQL DB! 🚀🚀")
}

func RetrieveDatabase() *sql.DB {
	return db
}
