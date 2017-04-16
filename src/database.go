package src

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"forcamp/conf"
	"log"
)

var(
	Connection *sql.DB // DataBase: "forcamp"
	CustomConnection *sql.DB // DataBase: custom
)

/*
Function builds MySQL source link for Database: "forcamp"
 */
func getMysqlSource() string{
	MysqlSource := conf.MYSQL_LOCAL_LOGIN + ":"
	MysqlSource += conf.MYSQL_LOCAL_PASSWORD + "@tcp("
	MysqlSource += conf.MYSQL_SERVER_ADDR + conf.MYSQL_SERVER_PORT + ")/"
	MysqlSource += conf.MYSQL_DB_MAIN
	return MysqlSource
}

/*
Function builds MySQL source link for Database: %name
 */
func getMysqlSource_Custom(name string) string{
	MysqlSource := conf.MYSQL_LOCAL_LOGIN + ":"
	MysqlSource += conf.MYSQL_LOCAL_PASSWORD + "@tcp("
	MysqlSource += conf.MYSQL_SERVER_ADDR + conf.MYSQL_SERVER_PORT + ")/"
	MysqlSource += name
	return MysqlSource
}

/*
Function opens an MySQL connection for Database: "ForCamp"
 */
func Connect() *sql.DB{
	Source := getMysqlSource()
	Connection, err := sql.Open("mysql", Source)
	if err != nil{
		log.Print(err)
	}
	Connection.SetMaxOpenConns(conf.MYSQL_MAX_USER_CONNECTIONS)
	Connection.SetMaxIdleConns(0)
	return Connection
}

/*
Function opens an MySQL connection for Database: %name
 */
func Connect_Custom(name string) *sql.DB{
	newConn, err := sql.Open("mysql", getMysqlSource_Custom(name))
	if err != nil{
		log.Print(err)
	}
	newConn.SetMaxOpenConns(conf.MYSQL_MAX_USER_CONNECTIONS)
	newConn.SetMaxIdleConns(0)
	return newConn
}