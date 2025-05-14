package initialize

import (
	"fmt"

	"github.com/TS0906/go-ecommerce-backend-api/global"
	"go.uber.org/zap"
)

func Run() {
	// Load configuration
	LoadConfig()
	m := global.Config.MySQL
	fmt.Print("Load configuration MySQL", m.Username, m.Password, m.Host, m.Port, m.Dbname)
	// Initialize logger
	InitLogger()
	global.Logger.Info("Config log ok!!", zap.String("ok", "success"))

	// Initialize Redis
	InitRedis()

	// Initialize MySQL
	InitMySQL()

	// Initialize router
	InitRouter()

	r := InitRouter()
	r.Run(":8002")
}
