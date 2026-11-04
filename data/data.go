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

func GetDialector() (*gorm.Dialector, error) {
	connString := os.Getenv("WORLD_SAYINGS_CONNECTION_STRING")
	if connString == "" {
		return nil, errors.New("Connection string not resolved. Ensure you create a .env file modeled after .env.template")
	}

	db := postgres.Open(connString)
	return &db, nil
}

func Migrate(dialector *gorm.Dialector) error {
	db, err := gorm.Open(*dialector, &gorm.Config{})

	if err != nil {
		return err
	}

	err = db.AutoMigrate(&Saying{}, &Tag{})

	if err != nil {
		return err
	}
	return nil
}
