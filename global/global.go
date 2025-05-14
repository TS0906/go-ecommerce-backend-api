package global

import (
	"github.com/TS0906/go-ecommerce-backend-api/pkg/logger"
	"github.com/TS0906/go-ecommerce-backend-api/pkg/setting"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Mdb    *gorm.DB
)
