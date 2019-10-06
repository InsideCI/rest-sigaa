package store

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres

	"github.com/InsideCI/nego/model"
	"github.com/joho/godotenv"
)

// var db *gorm.DB

// Store abstracts CRUD methods
type Store struct {
	db *gorm.DB
}

// NewStore creates and returns a database based on .env file.
func NewStore() *Store {

	err := godotenv.Load("database.env")
	if err != nil {
		log.Fatal("Error loading .env file.")
	}

	userName := os.Getenv("db_user")
	userPass := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbPort := os.Getenv("db_port")

	dbURI := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s  sslmode=disable", dbHost, dbPort, dbName, userName, userPass)

	db, err := gorm.Open("postgres", dbURI)
	if err != nil {
		log.Fatal(err)
		return nil

	}

	db.AutoMigrate(&model.Center{})
	return &Store{
		db: db,
	}
}

// GetCenters returns all centers available on UFPB SIGAA.
func (s *Store) GetCenters() []model.Center {
	var centers []model.Center

	fmt.Println("Finding centers.")
	s.db.Find(&centers)

	return centers
}
