package controllers

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
}

func NewController() *Controller {
	return &Controller{}
}

func (controller *Controller) HealthStatus(c *gin.Context) {
	empJson := `{
        "error" : false,
        "message" : "Success"
	}`

	var result map[string]interface{}
	err := json.Unmarshal([]byte(empJson), &result)
	if err != nil {
		log.Error("error while unmarshalling:", err)
	}

	c.JSON(http.StatusOK, result)
}
