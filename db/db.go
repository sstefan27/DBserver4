package db

import (
	"fmt"
	"go/problem4/config"

	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/jinzhu/gorm"
)

const dbErrorMessage = "Error connecting to DB"

var db *gorm.DB

func InitDatabase() {
	migrateConnection, err := migrate.New("file://db/migrate", config.GetConfig().Database.URL)
	if err != nil {
		fmt.Println("Error creating the message")
		return
	}
	version := config.GetConfig().Database.Version
	currentVersion, _, _ := migrateConnection.Version()
	if version != currentVersion {
		err = migrateConnection.Migrate(version)
		if err != nil {
			fmt.Println("Error creating the message")
			return
		}
	}
	migrateConnection.Close()
	db, err = gorm.Open("postgres", config.GetConfig().Database.URL)
	if err != nil {
		fmt.Println(dbErrorMessage)
	}
	fmt.Println(db)
	db.LogMode(true)
}

func GetDB() *gorm.DB {
	return db
}
