package data

import (
	"database/sql"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Saying struct {
	gorm.Model

	Original string
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

func SetUpDatabase() {
  db, err := gorm.Open(sqlite.Open("sayings.db"), &gorm.Config{})
  if err != nil {
    panic("failed to connect database")
  }

  db.AutoMigrate(&Saying{})
}

