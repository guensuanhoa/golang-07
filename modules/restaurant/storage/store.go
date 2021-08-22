package restaurantstorage

import (
	"gorm.io/gorm"
)

// muon khai niem Encapsulation cua OOP
// Co mot so ban de mot bien global: db => kho unit test
// {sqlStore} va {db} viet thuong => private
// Tang business khong duoc biet tuong minh tang storage
type sqlStore struct {
	db *gorm.DB
}

func NewSQLStore(db *gorm.DB) *sqlStore {
	return &sqlStore{db: db}
}
