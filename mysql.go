package gowaitfordeps

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

type MySQLConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
}

// WaitForMySQL waits for the MySQL database to be up and running
func WaitForMySQL(config MySQLConfig) {

	mysqlDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Name,
	)

	for {
		waitDb, err := sql.Open("mysql", mysqlDSN)
		if err == nil {
			err = waitDb.Ping()
			if err == nil {
				e := waitDb.Close()
				if e != nil {
					return
				}
				break
			}
		}

		log.Println("MySQL connection error, will retry:", err)
		time.Sleep(5 * time.Second)
	}

	log.Println("MySQL connection successful!")

}
