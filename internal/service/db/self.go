package db

import "gorm.io/gorm"

func provideDbMap() map[string]*gorm.DB {
	return make(map[string]*gorm.DB)
}
