package enviornment

import "gorm.io/gorm"

type Environment struct {
	gorm.Dialector
}
