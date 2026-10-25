package data

import (
	"database/sql"
	"errors"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Saying struct {
	gorm.Model

	Original    string
	Translation string
	Meaning     string
	WhenToUse   sql.NullString
	Example     sql.NullString

	Tags []*Tag `gorm:"many2many:Tag_Saying;"`
}

type Tag struct {
	gorm.Model

	Name string

	Sayings []*Saying `gorm:"many2many:Tag_Saying;"`
}

func SetUpDatabase() (*gorm.DB, error) {
	connString := os.Getenv("WORLD_SAYINGS_CONNECTION_STRING")
	if connString == "" {
		return nil, errors.New("Connection not resolved. Ensure you create a .env file modeled after .env.template")
	}

	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	migrateError := db.AutoMigrate(
		&Saying{},
		&Tag{},
	)

	if migrateError != nil {
		return nil, migrateError
	}

	return db, nil
}
