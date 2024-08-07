package repository

import (
	"basic-gin-app/internal/entity"
)

func CreateResource(ContentType *entity.Resource) (err error) {
	err = Repository().Db.Create(ContentType).Error
	if err != nil {
		return err
	}
	return nil
}

func GetResources(ContentType *[]entity.Resource) (err error) {
	err = Repository().Db.Find(ContentType).Error
	if err != nil {
		return err
	}
	return nil
}
