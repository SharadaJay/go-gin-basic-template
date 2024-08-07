package repository

import (
	"basic-gin-app/configs"

	"gorm.io/gorm"
)

type DatabaseRepo struct {
	Db *gorm.DB
}

func Repository() *DatabaseRepo {
	db := configs.GetDb()
	return &DatabaseRepo{Db: db}
}
