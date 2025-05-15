package initialize

import (
	"fmt"
	"time"

	"github.com/TS0906/go-ecommerce-backend-api/global"
	"github.com/TS0906/go-ecommerce-backend-api/internal/po"
	"go.uber.org/zap"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func checkErrPanic(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)
	}
}

func InitMySQL() {
	m := global.Config.MySQL
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	var s = fmt.Sprintf(dsn, m.Username, m.Password, m.Host, m.Port, m.Dbname)
	db, err := gorm.Open(mysql.Open(s), &gorm.Config{
		SkipDefaultTransaction: false,
	})
	checkErrPanic(err, "InitMySQL failed")
	global.Logger.Info("InitMySQL success")
	global.Mdb = db

	// set pool
	SetPool()
	// migrate tables
	migrateTables()
}

func SetPool() {
	m := global.Config.MySQL
	sqlDB, err := global.Mdb.DB()
	if err != nil {
		fmt.Printf("Mysql err: %s\n", err)
	}
	sqlDB.SetConnMaxIdleTime(time.Duration(m.MaxIdleConns))
	sqlDB.SetMaxOpenConns(m.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime))
}

func migrateTables() {
	err := global.Mdb.AutoMigrate(&po.User{}, &po.Role{})
	if err != nil {
		fmt.Printf("Migrate tables err: %s\n", err)
	}
}
