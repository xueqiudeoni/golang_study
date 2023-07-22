package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var db *sqlx.DB

func Init() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetInt("mysql.port"),
		viper.GetString("mysql.dbname"),
	)
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("connect db failed,err:", zap.Error(err))
		return
	}
	db.SetMaxIdleConns(viper.GetInt("mysql.max_idle_conns"))
	db.SetMaxOpenConns(viper.GetInt("mysql.max_open_conns"))
	return
}
func Close() {
	_ = db.Close()
}
