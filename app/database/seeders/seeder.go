package seeders

import (
	"github.com/idriscahyono/bookstore/app/database/fakers"
	"gorm.io/gorm"
)

type Seeder struct {
	Seeder interface{}
}

func BookSeders(db *gorm.DB) []Seeder {
	return []Seeder{
		{Seeder: fakers.BookFaker(db)},
	}
}

func DBSeed(db *gorm.DB) error {
	for _, s := range BookSeders(db) {
		err := db.Debug().Create(s.Seeder).Error
		if err != nil {
			return err
		}
	}

	return nil
}
