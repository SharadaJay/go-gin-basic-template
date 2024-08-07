package controllers

import (
	"basic-gin-app/internal/entity"
	"basic-gin-app/internal/services"
	"encoding/json"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (controller *Controller) CreateResource(c *gin.Context) {
	log.Debug("CreateResource method started")
	var resource entity.Resource

	// request validation
	if err := c.ShouldBindJSON(&resource); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := services.CreateResource(&resource)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	byteArray, errorObj := json.Marshal(resource)
	if errorObj != nil {
		log.Error("error while marshalling:", errorObj)
		c.AbortWithStatusJSON(http.StatusInternalServerError, "error while marshalling")
	}

	var resourceResponse entity.ResourceResponse
	if err := json.Unmarshal(byteArray, &resourceResponse); err != nil {
		log.Error("error while unmarshalling:", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, "error while unmarshalling")
	}

	log.Debug("CreateResource method ended")
	c.JSON(http.StatusCreated, resourceResponse)
}

func (controller *Controller) GetResources(c *gin.Context) {
	log.Debug("GetResources method started")
	var resources []entity.Resource
	err := services.GetResources(&resources)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	byteArray, errorObj := json.Marshal(resources)
	if errorObj != nil {
		log.Error("error while marshalling:", errorObj)
		c.AbortWithStatusJSON(http.StatusInternalServerError, "error while marshalling")
	}

	var resourceResponses []entity.ResourceResponse
	if err := json.Unmarshal(byteArray, &resourceResponses); err != nil {
		log.Error("error while unmarshalling:", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, "error while unmarshalling")
	}

	log.Debug("GetResources method ended")
	c.JSON(http.StatusOK, resourceResponses)
}
