package database

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

const (
	dbAlias       = "default"
	mysqlUser     = "root"
	mysqlPassword = "toor"
	mysqlHost     = "mysql"
	mysqlPort     = 3306
	mysqlDatabase = "watchme"
	mysqlCharset  = "utf8"
)

var (
	mysqlCon = fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s",
		mysqlUser,
		mysqlPassword,
		mysqlHost,
		mysqlPort,
		mysqlDatabase,
		mysqlCharset,
	)
)

type Database struct {
}

func NewDatabase() *Database {

	mysqlCon = fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s",
		mysqlUser,
		mysqlPassword,
		mysqlHost,
		mysqlPort,
		mysqlDatabase,
		mysqlCharset,
	)

	// register driver
	orm.RegisterDriver("mysql", orm.DRMySQL)

	// register model
	//orm.RegisterModel(new(User), new(Profile))

	// set timezone
	orm.DefaultTimeLoc = time.UTC

	// set default database
	orm.RegisterDataBase(dbAlias, "mysql", mysqlCon)


	fmt.Println("DB registered!")

	return &Database{

	}
}