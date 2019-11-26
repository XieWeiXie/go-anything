package mysql_operator

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/wuxiaoxiaoshen/go-anything/configs"
)

type (
	MySQLAction struct {
		DB *gorm.DB
	}
	mysqlSettings struct {
		port     string
		db       string
		user     string
		password string
		host     string
	}
	mysqlMap map[string]string
)

var (
	mysqlSetting       string
	DefaultMySQLAction MySQLAction
)

func MySQLInit() {
	m := configs.DefaultConfigs.LoadConfigs("mysql")
	a := m.(map[string]interface{})
	log.Println(fmt.Sprintf("Keys: MySQL: %#v", m))
	s := mysqlSettings{
		port:     a["port"].(string),
		db:       a["db"].(string),
		user:     a["user"].(string),
		password: a["passwd"].(string),
		host:     a["host"].(string),
	}
	mysqlSetting = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		s.user, s.password, s.host, s.port, s.db)
	DefaultMySQLAction.DB = newMysql()
}

func newMysql() *gorm.DB {

	db, e := gorm.Open("mysql", mysqlSetting)
	if e != nil {
		log.Println(fmt.Sprintf("mysql: newMysql: %s", e.Error()))
		log.Panic(fmt.Sprintf("mysql: newMysql : connect db error"))
		return nil
	}
	db.LogMode(true)
	db.DB().SetMaxIdleConns(3)
	return db
}

func (M MySQLAction) Close() {
	defer M.DB.Close()
}
