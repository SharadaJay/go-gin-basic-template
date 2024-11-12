package repository

import (
	"basic-gin-app/internal/entity"
	"errors"
	"gorm.io/gorm"
)

func CreateResource(input entity.ResourceRequest) (*entity.Resource, error) {
	resource := entity.Resource{
		Name:        input.Name,
		DisplayName: input.DisplayName,
		Type:        input.Type,
	}

	if err := Repository().Db.Create(&resource).Error; err != nil {
		return nil, err
	}

	return &resource, nil
}

func GetResourceByID(id uint) (*entity.Resource, error) {
	var resource entity.Resource

	if err := Repository().Db.First(&resource, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}

	return &resource, nil
}

func GetResources() ([]entity.Resource, error) {
	var resources []entity.Resource
	err := Repository().Db.Find(&resources).Error
	if err != nil {
		return nil, err
	}
	return resources, nil
}

func UpdateResource(id uint, newResourceData entity.Resource) (*entity.Resource, error) {
	var resource entity.Resource

	if err := Repository().Db.First(&resource, id).Error; err != nil {
		return nil, err
	}

	resource.Name = newResourceData.Name
	resource.DisplayName = newResourceData.DisplayName
	resource.Type = newResourceData.Type

	if err := Repository().Db.Save(&resource).Error; err != nil {
		return nil, err
	}

	return &resource, nil
}

func DeleteResource(id uint) error {
	if err := Repository().Db.Delete(&entity.Resource{}, id).Error; err != nil {
		return err
	}
	return nil
}
