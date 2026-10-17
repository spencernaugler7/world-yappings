package data

import (
	"database/sql"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Saying struct {
	gorm.Model

	ID uint
	Original string
	Meaning string
	WhenToUse sql.NullString
	Example sql.NullString
	Created time.Time
	LastUpdated sql.NullTime

	Tags []*Tag `gorm:"many2many:Tag_Saying;"`
}

type Tag struct {
	gorm.Model

	ID uint
	Name string
	Created time.Time
	LastUpdated sql.NullTime

	Sayings []*Saying `gorm:"many2many:Tag_Saying;"`
}

func SetUpDatabase() {
  db, err := gorm.Open(sqlite.Open("sayings.db"), &gorm.Config{})
  if err != nil {
    panic("failed to connect database")
  }

  db.AutoMigrate(&Saying{})
}

