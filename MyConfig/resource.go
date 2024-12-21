package MyConfig

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

var Logger *zap.SugaredLogger
var DB *gorm.DB
var OutTime = 30 * 60 * time.Second
