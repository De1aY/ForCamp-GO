package src

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"forcamp/conf"
	"log"
)

func getMysqlSource() string{
	MysqlSource := conf.MYSQL_LOGIN + ":"
	MysqlSource += conf.MYSQL_PASSWORD + "@tcp("
	MysqlSource += conf.MYSQL_SERVER_ADDR + conf.MYSQL_SERVER_PORT + ")/"
	MysqlSource += conf.MYSQL_DB_MAIN
	return MysqlSource
}

func Connect() *sql.DB{
	Source := getMysqlSource()
	Connection, err := sql.Open("mysql", Source)
	if err != nil{
		log.Fatal(err)
	}
	return Connection
}