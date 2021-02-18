package databases

import (
	"fmt"

	"github.com/JGurus/template-initial-api/databases/interfaces"
	"github.com/JGurus/template-initial-api/databases/mysql"
)

//Factory select the database engine implementation
func Factory(engine string) interfaces.User {
	switch engine {
	case "mysql":
		return mysql.NewUserImpl(GetMYSQLConnection())
	default:
		fmt.Printf("%s database engine is not implemented", engine)
	}
	return nil
}
