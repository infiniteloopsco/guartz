package models

import (
	"database/sql"
	"os"

	"github.com/infiniteloopsco/guartz/utils"

	"github.com/jinzhu/gorm"
)

//Gdb connection
var Gdb *gorm.DB

//InitDB connection
func InitDB() {
	//open db
	utils.Log.Info("*** INIT DB ***")
	connString := os.Getenv("MYSQL_DB")
	db, err := gorm.Open("mysql", connString)
	if err != nil {
		utils.Log.Panic(err)
	}
	db.DB().Ping()
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	//Migrations
	db.AutoMigrate(&Task{})
	db.AutoMigrate(&Execution{})

	//Add FK
	db.Model(&Execution{}).AddForeignKey("task_id", "tasks(id)", "CASCADE", "CASCADE")

	Gdb = &db
}

//InTx executes function in a transaction
func InTx(f func(*gorm.DB) bool) {
	utils.Log.Info("***INIT TRANSACTION***")
	txn := Gdb.Begin()
	if txn.Error != nil {
		utils.Log.Panic(txn.Error)
	}
	if f(txn) == true {
		utils.Log.Info("***TRANSACTION COMMITED***")
		txn.Commit()
	} else {
		utils.Log.Info("***TRANSACTION ROLLBACK***")
		txn.Rollback()
	}
	if err := txn.Error; err != nil && err != sql.ErrTxDone {
		utils.Log.Panic(err)
	}
}
