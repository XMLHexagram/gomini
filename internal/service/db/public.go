package db

import (
	"errors"
	"gorm.io/gorm"
)

func ProvideDb(Name string) (*gorm.DB, error) {
	db, ok := dbService.DbMap[Name]
	if !ok {
		return nil, errors.New("not found")
	}
	return db, nil
}
