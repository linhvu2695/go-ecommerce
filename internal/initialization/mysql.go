package initialization

import (
	"database/sql"
	"fmt"
	"go-ecommerce/global"

	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

func InitMySql() {
	m := global.Config.Mysql

	dsnTemplate := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf(dsnTemplate, m.Username, m.Password, m.Host, m.Port, m.DbName)

	db, err := sql.Open("mysql", dsn)
	checkErrorPanic(err, "Failed to init MySQL")
	global.Logger.Info("Initialize MySQL successfully")

	global.Db = db

	SetPool()
	migrateTables()
}

func SetPool() {
	// m := global.Config.Mysql

	// db, err := global.Db.DB()
	// if err != nil {
	// 	checkErrorPanic(err, "Failed to set pool for MySQL")
	// }

	// db.SetConnMaxIdleTime(time.Duration(m.MaxIdleConns))
	// db.SetMaxOpenConns(m.MaxOpenConns)
	// db.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime))
}

func migrateTables() {
	// err := global.Db.AutoMigrate(
	// 	&po.User{},
	// 	&po.Role{},
	// )

	// if err != nil {
	// 	global.Logger.Error("Migrating tables failed", zap.Error(err))
	// }
}

func checkErrorPanic(err error, errMsg string) {
	if err != nil {
		global.Logger.Error(errMsg, zap.Error(err))
		panic(err)
	}
}
