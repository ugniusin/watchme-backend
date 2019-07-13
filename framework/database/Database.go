package database

import (
	"fmt"
	config2 "github.com/ugniusin/watchme/framework/config"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

type Database struct {
	configuration *config2.Configuration
}

func NewDatabase(configuration *config2.Configuration) *Database {
	return &Database{
		configuration: configuration,
	}
}

func (database *Database) BootstrapDatabase() {

	fmt.Println(database.configuration.Database)

	mysqlCon := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=%s",
		database.configuration.Database["mysqlUser"],
		database.configuration.Database["mysqlPassword"],
		database.configuration.Database["mysqlHost"],
		database.configuration.Database["mysqlPort"],
		database.configuration.Database["mysqlDatabase"],
		database.configuration.Database["mysqlCharset"],
	)

	// register driver
	orm.RegisterDriver("mysql", orm.DRMySQL)

	// register model
	//orm.RegisterModel(new(User), new(Profile))

	// set timezone
	orm.DefaultTimeLoc = time.UTC

	// set default database
	orm.RegisterDataBase(database.configuration.Database["dbAlias"], "mysql", mysqlCon)


	fmt.Println("DB registered!")
}
