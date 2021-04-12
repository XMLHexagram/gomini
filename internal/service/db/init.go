package db

import (
	"github.com/lmx-Hexagram/gemini-generator/internal/service/config"
	"github.com/lmx-Hexagram/gemini-generator/internal/service/log"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbService *service

type service struct {
	DbMap       map[string]*gorm.DB
	DbConfigMap config.DbMap
}

func Init() {
	s := InitDep()

	for k, v := range s.DbConfigMap {
		dialector := mysql.Open(v.DSN)
		db, err := gorm.Open(dialector, &gorm.Config{})
		if err != nil {
			log.Panic("can not open db", zap.Error(err))
		}
		s.DbMap[k] = db
	}

	dbService = s

	log.Info("dbService init successfully")
}
