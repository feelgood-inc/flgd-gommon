package utils

import (
	"github.com/feelgood-inc/flgd-gommon/models"
	"gorm.io/gorm"
	"math"
)

func Paginate(value interface{}, pagination *models.Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	db.Find(value).Count(&totalRows)

	rows := int(totalRows)
	pagination.TotalItems = &rows
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.ItemsPerPage)))
	pagination.TotalPages = &totalPages

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset((pagination.Page - 1) * pagination.ItemsPerPage).Limit(pagination.ItemsPerPage)
	}
}
