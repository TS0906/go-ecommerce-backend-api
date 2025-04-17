package initialize

import (
	"fmt"

	"github.com/TS0906/go-ecommerce-backend-api/global"
)

func Run() {
	// Load configuration
	LoadConfig()
	m := global.Config.MySQL
	fmt.Print("Load configuration MySQL", m.Username, m.Password, m.Host, m.Port, m.Dbname)
	// Initialize logger
	InitLogger()

	// Initialize Redis
	InitRedis()

	// Initialize MySQL
	InitMySQL()

	// Initialize router
	InitRouter()

	r := InitRouter()
	r.Run(":8002")
}
