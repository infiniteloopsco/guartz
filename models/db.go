package models

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
)

//Gdb connection
var Gdb *gorm.DB

//InitDB connection
func InitDB() {
	//open db
	fmt.Println("*** INIT DB ***")
	connString := os.Getenv("MYSQL_DB")
	db, err := gorm.Open("mysql", connString)
	if err != nil {
		fmt.Println("Unable to connect to the database")
		panic(err)
	}
	db.DB().Ping()
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	//Migrations
	db.AutoMigrate(&Task{})
	db.AutoMigrate(&Execution{})

	//Add FK
	db.Model(&Execution{}).AddForeignKey("task_id", "tasks(id)", "RESTRICT", "RESTRICT")

	Gdb = &db
}

//InTx executes function in a transaction
func InTx(f func(*gorm.DB) bool) {
	txn := Gdb.Begin()
	if txn.Error != nil {
		panic(txn.Error)
	}
	if f(txn) == true {
		txn.Commit()
	} else {
		txn.Rollback()
	}
	if err := txn.Error; err != nil && err != sql.ErrTxDone {
		panic(err)
	}
}
