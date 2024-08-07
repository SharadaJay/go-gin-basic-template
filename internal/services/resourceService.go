package services

import (
	"basic-gin-app/internal/entity"
	"basic-gin-app/internal/repository"
	log "github.com/sirupsen/logrus"
)

// CreateResource create a Resource
func CreateResource(resource *entity.Resource) (err error) {
	err = repository.CreateResource(resource)
	if err != nil {
		log.Error("error while creating resource:", err)
	}
	return nil
}

// GetResources get all Resources
func GetResources(resources *[]entity.Resource) (err error) {
	err = repository.GetResources(resources)
	if err != nil {
		log.Error("error while creating resource:", err)
	}
	return nil
}
