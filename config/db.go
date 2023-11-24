package config

import (
	"echo-framework/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "docker-postgres"
	password = "docker-postgres"
	dbname   = "postgresql"
)

var (
	db  *gorm.DB
	err error
)

// func Connect() (*sql.DB, error) {
// 	psql := fmt.Sprintf(`
// 		host=%s
// 		port=%d
// 		user=%s`+`
// 		password=%s
// 		dbname=%s
// 		sslmode=disable`, host, port, user, password, dbname)

// 	// db, err = sql.Open("postgres", psql)  --koneksi native

// 	// defer db.Close()

// 	err = db.Ping()
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println("Succesfully connect to db")

// 	return db, nil
// }

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", host, user, password, dbname, port)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to database :", err)
	}

	db.AutoMigrate(models.User{}, models.Product{}, models.Employee{})
}

func GetDB() *gorm.DB {
	return db
}
