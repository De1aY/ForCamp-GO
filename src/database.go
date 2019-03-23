package src

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"wplay/conf"
	"log"
)

var(
	Connection *sql.DB // DataBase: "wplay"
	CustomConnection *sql.DB // DataBase: custom
)

/*
Function builds MySQL source link for Database: "wplay"
 */
func getMysqlSource() string{
	MysqlSource := conf.MYSQL_LOGIN + ":"
	MysqlSource += conf.MYSQL_PASSWORD + "@tcp("
	MysqlSource += conf.MYSQL_SERVER_ADDR + conf.MYSQL_SERVER_PORT + ")/"
	MysqlSource += conf.MYSQL_DB_MAIN
	return MysqlSource
}

/*
Function builds MySQL source link for Database: %name
 */
func getMysqlSource_Custom(name string) string{
	MysqlSource := conf.MYSQL_LOGIN + ":"
	MysqlSource += conf.MYSQL_PASSWORD + "@tcp("
	MysqlSource += conf.MYSQL_SERVER_ADDR + conf.MYSQL_SERVER_PORT + ")/"
	MysqlSource += name
	return MysqlSource
}

func getMysqlSource_Admin() string{
	MysqlSource := conf.MYSQL_LOGIN + ":"
	MysqlSource += conf.MYSQL_PASSWORD + "@tcp("
	MysqlSource += conf.MYSQL_SERVER_ADDR + conf.MYSQL_SERVER_PORT + ")/"
	return MysqlSource
}

/*
Function opens an MySQL connection for Database: "wplay"
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
func Connect_Custom(organizationName string) *sql.DB{
	newConn, err := sql.Open("mysql", getMysqlSource_Custom(organizationName))
	if err != nil{
		log.Print(err)
	}
	newConn.SetMaxOpenConns(conf.MYSQL_MAX_USER_CONNECTIONS)
	newConn.SetMaxIdleConns(0)
	return newConn
}

func Connect_Admin() *sql.DB{
	newConn, err := sql.Open("mysql", getMysqlSource_Admin())
	if err != nil{
		log.Print(err)
	}
	newConn.SetMaxOpenConns(conf.MYSQL_MAX_USER_CONNECTIONS)
	newConn.SetMaxIdleConns(0)
	return newConn
}