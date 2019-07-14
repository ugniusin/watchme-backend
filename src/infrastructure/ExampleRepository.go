package infrastructure

import (
	"fmt"
	"log"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)


type User struct {
	ID      int
	Name    string   `orm:"size(100)"`
	Profile *Profile `orm:"rel(one)"` // OneToOne relation
}

type Profile struct {
	ID   int
	Age  int16
	User *User `orm:"reverse(one)"` // Reverse relationship
}

const (
	dbAlias       = "default"
	mysqlUser     = "root"
	mysqlPassword = "toor"
	mysqlHost     = "mysql"
	mysqlPort     = 3306
	mysqlDatabase = "watchme-backend"
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

func init() {
	// register driver
	orm.RegisterDriver("mysql", orm.DRMySQL)

	// register model
	orm.RegisterModel(new(User), new(Profile))

	// set timezone
	orm.DefaultTimeLoc = time.UTC

	// set default database
	orm.RegisterDataBase(dbAlias, "mysql", mysqlCon)
}

func Main() {
	// Print SQL
	orm.Debug = true

	force := true   // Drop table and re-create.
	verbose := true // Print log
	// generate Tables
	if err := orm.RunSyncdb(dbAlias, force, verbose); err != nil {
		log.Println(err)
	}

	o := orm.NewOrm()
	o.Using("default")

	profileA := new(Profile)
	profileA.Age = 30

	profileB := new(Profile)
	profileB.Age = 40

	profileC := new(Profile)
	profileC.Age = 50

	profiles := []*Profile{profileB, profileC}

	user := new(User)
	user.Profile = profileA
	user.Name = "alpha"

	var res int64
	var err error

	// Insert
	if res, err = o.Insert(profileA); err != nil {
		log.Println(err)
	}
	log.Printf("inserted: %d row", res)

	if res, err = o.InsertMulti(10, profiles); err != nil {
		log.Println(err)
	}
	log.Printf("inserted: %d row", res)

	if res, err = o.Insert(user); err != nil {
		log.Println(err)
	}
	log.Printf("inserted: %d row", res)

	// Update
	updateTarget := User{ID: 1}
	if o.Read(&updateTarget) == nil {
		updateTarget.Name = "ALPHA"
		if num, err := o.Update(&updateTarget); err != nil {
			log.Println(err)
		} else {
			log.Printf("updated %d rows", num)
		}
	}

	// Read
	readTarget := User{ID: 1}
	err = o.Read(&readTarget)
	if err == orm.ErrNoRows {
		log.Println("No result found.")
	} else if err == orm.ErrMissPK {
		log.Println("No primary key found.")
	} else {
		log.Println(readTarget.ID, readTarget.Name)
	}

	// Raw Query
	var params []orm.Params
	num, err := o.Raw("SELECT * FROM profile WHERE age > ?", 30).Values(&params)
	if err != nil {
		log.Println(err)
	}
	log.Printf("selectd %d rows", num)
	if num > 0 {
		log.Printf("result %v", params)
	}
}