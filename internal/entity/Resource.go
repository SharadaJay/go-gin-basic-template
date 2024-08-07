package entity

import (
	"gorm.io/gorm"
)

type Resource struct {
	gorm.Model
	Name        string `gorm:"varchar(50)" json:"name" binding:"required"`
	DisplayName string `gorm:"varchar(50)" json:"display_name"`
	Type        string `gorm:"varchar(50)" json:"type" binding:"required"`
}

type ResourceRequest struct {
	Name        string `json:"name" binding:"required"`
	DisplayName string `json:"display_name"`
	Type        string `json:"type" binding:"required"`
}

type ResourceResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	Type        string `json:"type"`
}
