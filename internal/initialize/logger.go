package initialize

import (
	"github.com/TS0906/go-ecommerce-backend-api/global"
	"github.com/TS0906/go-ecommerce-backend-api/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
