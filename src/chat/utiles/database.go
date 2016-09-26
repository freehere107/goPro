package utiles

import (
	_ "github.com/go-sql-driver/mysql" // only use form gorm
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var (
	DB *gorm.DB
)

func init() {
	dbPath := viper.GetString("mysql.host")
	dbUser := viper.GetString("mysql.user")
	dbPass := viper.GetString("mysql.pass")
	dbName := viper.GetString("mysql.dbName")
	DB = initMysql(dbPath, dbUser, dbPass, dbName)
}

func initMysql(host, user, pass, name string) *gorm.DB {
	tdb, err := gorm.Open("mysql", user+":"+pass+"@tcp("+host+")/"+name+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	tdb.DB().SetMaxIdleConns(10)
	tdb.DB().SetMaxOpenConns(100)
	return tdb
}
