package config

import (
	"fmt"
	"os"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB is the main database connection
var (
	connection *gorm.DB
	once       sync.Once
)

// pgOptions is a struct to store the database connection options
type pgOptions struct {
	Host     string
	User     string
	Password string
	Dbname   string
	Port     string
}

// buildDSN builds the Data Source Name (DSN) string
func (p *pgOptions) buildDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", p.Host, p.User, p.Password, p.Dbname, p.Port)
}

func NewPostgresConnection() *gorm.DB {
	once.Do(func() {
		connection = getConnection()
	})
	return connection
}

// MainDbConnection is the function that connects to the database
func MainDbConnection() {
	loadEnv()

	dsn := environmentDSN()

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic(err)
	}

	DB = db

	fmt.Println("Successful connection to the main database")
}

// environmentDSN builds the Data Source Name (DSN) string from environment variables
func environmentDSN() string {
	options := pgOptions{
		Host:     os.Getenv("MAIN_DB_HOST"),
		User:     os.Getenv("MAIN_DB_USER"),
		Password: os.Getenv("MAIN_DB_PASSWORD"),
		Dbname:   os.Getenv("MAIN_DB_NAME"),
		Port:     os.Getenv("MAIN_DB_PORT"),
	}

	return options.buildDSN()
}
