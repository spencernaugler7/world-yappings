package data

import (
	"database/sql"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DataError struct {
	Message string
}

func (m *DataError) Error() string {
	return m.Message
}

type Saying struct {
	gorm.Model

	Original string
	Translation string
	Meaning string
	WhenToUse sql.NullString
	Example sql.NullString

	Tags []*Tag `gorm:"many2many:Tag_Saying;"`
}

type Tag struct {
	gorm.Model

	Name string

	Sayings []*Saying `gorm:"many2many:Tag_Saying;"`
}

func SetUpDatabase() error {
	db, err := gorm.Open(sqlite.Open("sayings.db"), &gorm.Config{})
	
	if err != nil {
		return err
	}

	connString := os.Getenv("WORLD_SAYINGS_CONNECTION_STRING")
	if connString == "" {
		return &DataError{ "connection string not found in enviornment variable. Ensure you create a .env file (template is .env.template)" }
	}

	migrateError := db.AutoMigrate(
		&Saying{},
		&Tag{},
	)

	return migrateError
}

