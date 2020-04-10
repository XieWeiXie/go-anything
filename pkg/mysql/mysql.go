package mysql_operator

import (
	"database/sql"
	"fmt"
	"log"
	"os"

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
)

var (
	mysqlSetting       string
	mysqlSettingORI    string
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
	if os.Getenv(configs.MYSQL_HOST) != "" {
		s.host = os.Getenv(configs.MYSQL_HOST)

	}
	//s.host = "119.3.198.221"
	if os.Getenv(configs.MYSQL_PORT) != "" {
		s.port = os.Getenv(configs.MYSQL_PORT)

	}
	//s.port = "30000"
	mysqlSetting = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		s.user, s.password, s.host, s.port, s.db)

	log.Println(mysqlSetting)
	mysqlSettingORI = mysqlSetting
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

type MysqlActionWithORI struct {
	DB *sql.DB
}

func (m MysqlActionWithORI) Close() {
	m.DB.Close()
}

var (
	MySQLActionWithORI MysqlActionWithORI
)

func MysqlInitORI() {
	db, e := sql.Open("mysql", mysqlSettingORI)
	if e != nil {
		log.Panicln(e)
		panic(e)
	}
	if e := db.Ping(); e != nil {
		log.Panicln(e)
		panic(e)
	}
	db.SetMaxIdleConns(3)
	db.SetMaxOpenConns(3)

	MySQLActionWithORI.DB = db
}

var DefaultMySQLActionWithORI = MySQLActionWithORI
