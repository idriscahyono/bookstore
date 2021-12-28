package fakers

import (
	"github.com/bxcodec/faker/v3"
	"github.com/idriscahyono/bookstore/app/models"
	"gorm.io/gorm"
)

func BookFaker(*gorm.DB) *models.Book {
	return &models.Book{
		Title:  faker.TitleMale(),
		Author: faker.Word(),
	}
}
