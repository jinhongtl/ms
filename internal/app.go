package internal

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var AppConfig Config

var Db *gorm.DB

var Logger *logrus.Logger
