package initialize

import (
	"awesomeProject/global"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
	"time"
)

func InitMysql() {
	m := global.Config.Mysql

	dsn := "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	connectString := fmt.Sprintf(dsn, m.UserName, m.Password, m.Host, m.Port, m.DbName)

	db, err := sql.Open("mysql", connectString)
	checkError(err, "InitMysql error")
	global.Logger.Info("connect ok")

	global.Mdbcc = db

	SetPool()
	//migrateTable()

}
func checkError(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)
	}

}

func SetPool() {
	m := global.Config.Mysql

	// Lấy đối tượng *sql.DB từ *gorm.DB
	sqlDb := global.Mdbcc
	// Thiết lập các tham số cho pool kết nối
	sqlDb.SetConnMaxLifetime(time.Duration(m.ConnMaxLifeTime) * time.Second)
	sqlDb.SetMaxOpenConns(m.MaxOpenCons)
	sqlDb.SetMaxIdleConns(m.MaxIdleCons)

	// Log thông tin cấu hình pool
	global.Logger.Info("Connection pool configured successfully")
}
