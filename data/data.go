package data

import (
	"database/sql"
	"os"

	"gorm.io/driver/postgres"
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

func SetUpDatabase() error {
	connString := os.Getenv("WORLD_SAYINGS_CONNECTION_STRING")
	if connString == "" {
		return &DataError{"Connection not resolved. Ensure you create a .env file modeled after .env.template"}
	}

	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{})

	if err != nil {
		return err
	}

	migrateError := db.AutoMigrate(
		&Saying{},
		&Tag{},
	)

	return migrateError
}
