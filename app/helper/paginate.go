package helper

import (
	"gorm.io/gorm"
)

const (
	defaultLimit = 10
	maxLimit     = 10000
)

func Paginate(limit, page int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		adjustedPage := adjustPage(page)
		adjustedLimit := adjustLimit(limit)
		offset := calculateOffset(adjustedPage, adjustedLimit)

		return db.Offset(offset).Limit(adjustedLimit)
	}
}

func adjustPage(page int) int {
	if page <= 0 {
		return 1
	}
	return page
}

func adjustLimit(limit int) int {
	if limit <= 0 {
		return defaultLimit
	}
	if limit > maxLimit {
		return maxLimit
	}
	return limit
}

func calculateOffset(page, limit int) int {
	return (page - 1) * limit
}
