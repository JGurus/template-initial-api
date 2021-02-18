package databases

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	"github.com/JGurus/template-initial-api/config"
	"github.com/JGurus/template-initial-api/models"

	//mysql driver
	_ "github.com/go-sql-driver/mysql"
)

var (
	once sync.Once
	db   *sql.DB
)

//GetMYSQLConnection .
func GetMYSQLConnection() *sql.DB {
	once.Do(mysqlConnection)
	return db
}

func mysqlConnection() {
	var err error
	configDB := getConfig()
	uri := fmt.Sprintf("%s:%s@/%s?tls=false&autocommit=true&parseTime=true", configDB.User, configDB.Password, configDB.Database)
	db, err = sql.Open(configDB.Engine, uri)
	if err != nil {
		log.Fatalf("No se pudo conectar a la bd: %v", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("No se pudo hacer ping: %v", err)
	}
	fmt.Println("Conectado a la BD")
}

func getConfig() models.ConfigDB {
	var err error
	c := models.ConfigDB{}
	c, err = config.GetConfigDB()
	if err != nil {
		log.Fatal("no se encontró el archivo")
		return c
	}
	return c
}
