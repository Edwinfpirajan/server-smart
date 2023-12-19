package src

import (
	"fmt"
	"sync"

	"github.com/Edwinfpirajan/server-smart.git/config"
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
		connection = MainDbConnection()
	})
	return connection
}

// MainDbConnection is the function that connects to the database
func MainDbConnection() *gorm.DB {
	config.Environments()

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

	fmt.Println("Successful connection to the main database")

	return db
}

// environmentDSN builds the Data Source Name (DSN) string from environment variables
func environmentDSN() string {
	options := pgOptions{
		Host:     config.Cfg.MainDb.Host,
		User:     config.Cfg.MainDb.User,
		Password: config.Cfg.MainDb.Password,
		Dbname:   config.Cfg.MainDb.DbName,
		Port:     fmt.Sprintf("%d", config.Cfg.MainDb.Port),
	}

	return options.buildDSN()
}
