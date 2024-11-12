package services

import (
	"basic-gin-app/internal/entity"
	"basic-gin-app/internal/repository"
	log "github.com/sirupsen/logrus"
)

// CreateResource create a Resource
func CreateResource(input entity.ResourceRequest) (*entity.ResourceResponse, error) {

	resource, err := repository.CreateResource(input)
	if err != nil {
		log.Error("error while creating resource:", err)
		return nil, err
	}

	// Convert the saved resource into a ResourceResponse
	resourceResponse := &entity.ResourceResponse{
		ID:          resource.ID,
		Name:        resource.Name,
		DisplayName: resource.DisplayName,
		Type:        resource.Type,
	}

	return resourceResponse, nil
}

// GetResourceByID get one Resource by ID
func GetResourceByID(id uint) (*entity.ResourceResponse, error) {
	resource, err := repository.GetResourceByID(id)
	if err != nil {
		log.Error("error while retrieving resource:", err)
		return nil, err
	}

	response := &entity.ResourceResponse{
		ID:          resource.ID,
		Name:        resource.Name,
		DisplayName: resource.DisplayName,
		Type:        resource.Type,
	}

	return response, nil
}

// GetResources get all Resources
func GetResources() ([]entity.ResourceResponse, error) {

	resources, err := repository.GetResources()
	if err != nil {
		log.Error("error while fetching resources:", err)
		return nil, err
	}

	var resourceResponses []entity.ResourceResponse
	for _, resource := range resources {
		resourceResponses = append(resourceResponses, entity.ResourceResponse{
			ID:          resource.ID,
			Name:        resource.Name,
			DisplayName: resource.DisplayName,
			Type:        resource.Type,
		})
	}

	return resourceResponses, nil
}

// UpdateResource updates a resource by ID
func UpdateResource(id uint, resourceRequest entity.ResourceRequest) (*entity.ResourceResponse, error) {

	resource := entity.Resource{
		Name:        resourceRequest.Name,
		DisplayName: resourceRequest.DisplayName,
		Type:        resourceRequest.Type,
	}

	updatedResource, err := repository.UpdateResource(id, resource)
	if err != nil {
		log.Error("error while updating resource:", err)
		return nil, err
	}

	response := &entity.ResourceResponse{
		ID:          updatedResource.ID,
		Name:        updatedResource.Name,
		DisplayName: updatedResource.DisplayName,
		Type:        updatedResource.Type,
	}

	return response, nil
}

// DeleteResource delete a resource by ID
func DeleteResource(id uint) error {
	err := repository.DeleteResource(id)
	if err != nil {
		log.Error("error while deleting resource:", err)
	}
	return err
}
