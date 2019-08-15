package dao


import (
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql" //mysql driver
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// func init() {
// var err error
// db, err = gorm.Open("mysql", dsn)

// if err != nil {
// 	log.Println(err)
// }

// // SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
// db.DB().SetMaxIdleConns(10)

// // SetMaxOpenConns sets the maximum number of open connections to the database.
// db.DB().SetMaxOpenConns(100)
// }

const dsn string = "test:test123@tcp(localhost:3306)/szhouse?charset=utf8&parseTime=True&loc=Local"

/*
Insert new record
*/
func Insert(b ProjectBrief) bool {

	e := GetDBConn().Table("project_brief").Create(&b).Error
	if e != nil {
		log.Println(e)
		return false
	}
	return true

	// flag := db.Table("project_brief").NewRecord(b)
	// if flag {
	// 	log.Println("insert failed")
	// }
	// return !flag

}

//GetDBConn get database connection
func GetDBConn() *gorm.DB {

	if db == nil {
		var mu sync.Mutex
		mu.Lock()
		defer mu.Unlock()

		log.Print("start connect to database")
		var err error
		db, err = gorm.Open("mysql", dsn)

		if err != nil {
			log.Println(err)
		}

		// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
		db.DB().SetMaxIdleConns(10)

		// SetMaxOpenConns sets the maximum number of open connections to the database.
		db.DB().SetMaxOpenConns(10)

		return db
	}
	return db

}
