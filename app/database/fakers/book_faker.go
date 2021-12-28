package fakers

import (
	"gin-bookstore/app/models"
	"github.com/bxcodec/faker/v3"
	"gorm.io/gorm"
)

func BookFaker(*gorm.DB) *models.Book {
	return &models.Book{
		Title:  faker.TitleMale(),
		Author: faker.Word(),
	}
}
