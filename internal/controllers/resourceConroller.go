package controllers

import (
	"basic-gin-app/internal/entity"
	"basic-gin-app/internal/services"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func (controller *Controller) CreateResource(c *gin.Context) {
	log.Debug("CreateResource method started")
	var resourceRequest entity.ResourceRequest

	if err := c.ShouldBindJSON(&resourceRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resourceResponse, err := services.CreateResource(resourceRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Debug("CreateResource method ended")
	c.JSON(http.StatusCreated, resourceResponse)
}

func (controller *Controller) GetResourceByID(c *gin.Context) {
	log.Debug("GetResourceByID method started")

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid resource ID"})
		return
	}

	resource, err := services.GetResourceByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	log.Debug("GetResourceByID method ended")
	c.JSON(http.StatusOK, resource)
}

func (controller *Controller) GetResources(c *gin.Context) {
	log.Debug("GetResources method started")

	resourceResponses, err := services.GetResources()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	log.Debug("GetResources method ended")
	c.JSON(http.StatusOK, resourceResponses)
}

func (controller *Controller) UpdateResource(c *gin.Context) {
	log.Debug("UpdateResource method started")

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid resource ID"})
		return
	}

	var resourceRequest entity.ResourceRequest
	if err := c.ShouldBindJSON(&resourceRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedResource, err := services.UpdateResource(uint(id), resourceRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Debug("UpdateResource method ended")
	c.JSON(http.StatusOK, updatedResource)
}

func (controller *Controller) DeleteResource(c *gin.Context) {
	log.Debug("DeleteResource method started")

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid resource ID"})
		return
	}

	err = services.DeleteResource(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	log.Debug("DeleteResource method ended")
	c.JSON(http.StatusNoContent, gin.H{"message": "Resource deleted successfully"})
}
