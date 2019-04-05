/*
	Copyright: "NullTeam", 2016 - 2019
	Author: Nikita Ivanov <de1ay@nullteam.info>
*/
package src

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"wplay/conf"
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
	MysqlSource := conf.MysqlLogin + ":"
	MysqlSource += conf.MysqlPassword + "@tcp("
	MysqlSource += conf.MysqlServerAddr + conf.MysqlServerPort + ")/"
	MysqlSource += conf.MySqlDbMain
	return MysqlSource
}

/*
Function builds MySQL source link for Database: %name
 */
func getMysqlSource_Custom(name string) string{
	MysqlSource := conf.MysqlLogin + ":"
	MysqlSource += conf.MysqlPassword + "@tcp("
	MysqlSource += conf.MysqlServerAddr + conf.MysqlServerPort + ")/"
	MysqlSource += name
	return MysqlSource
}

func getMysqlSource_Admin() string{
	MysqlSource := conf.MysqlLogin + ":"
	MysqlSource += conf.MysqlPassword + "@tcp("
	MysqlSource += conf.MysqlServerAddr + conf.MysqlServerPort + ")/"
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
	Connection.SetMaxOpenConns(conf.MysqlMaxUserConnections)
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
	newConn.SetMaxOpenConns(conf.MysqlMaxUserConnections)
	newConn.SetMaxIdleConns(0)
	return newConn
}

func Connect_Admin() *sql.DB{
	newConn, err := sql.Open("mysql", getMysqlSource_Admin())
	if err != nil{
		log.Print(err)
	}
	newConn.SetMaxOpenConns(conf.MysqlMaxUserConnections)
	newConn.SetMaxIdleConns(0)
	return newConn
}
