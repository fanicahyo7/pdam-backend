package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"pdam/repository"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	//cek env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbusername := os.Getenv("DB_USERNAME")
	if dbusername == "" {
		log.Fatal("Username is required!")
	}

	dbpassword := os.Getenv("DB_PASSWORD")

	dbhost := os.Getenv("DB_HOST")
	if dbusername == "" {
		log.Fatal("Host is required!")
	}

	dbport := os.Getenv("DB_PORT")
	if dbusername == "" {
		log.Fatal("Port is required!")
	}

	dbdatabase := os.Getenv("DB_DATABASE")
	if dbusername == "" {
		log.Fatal("Database is required!")
	}

	dsn := dbusername + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbdatabase + "?charset=utf8mb4&parseTime=True&loc=Local"

	//koneksi ke database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println(rand.Int())
	fmt.Println(rand.Int())
	fmt.Println(rand.Int())

	grouprepository := repository.NewGroupRepository(db)

	fmt.Println(grouprepository.FindGroupByID(0))

}
