package setup

import (
	"{{ .AppName }}/db"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return nil
}

func All() (err error) {
	if err := db.Setup(); err != nil {
		return err
	}

	if Migrate(db.GetDB()) != nil {
		return err
	}

	return nil
}

func Down() (err error) {
	db.CloseDB()
	return nil
}
