package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

var DbInit = DbWorkerConnect{}
type DbWorkerConnect struct {
	Conn *sqlx.DB
}
func main () {

	db, err := sqlx.Connect("mysql", "root:Pere@123@tcp(127.0.0.1:3306)/golangpoc")
	if err != nil {
		log.Fatalln(err)
	}else {
		log.Println("Data successfully updated")
	}
	db.SetMaxIdleConns(0)
	DbInit.Conn = db

	db.Query("insert into golangpoc.user_registration(regid,FN,LN,mobno,emailid,userid,passwrd) values(5,'Mayur','Khairnar',95459,'mayurkhairnar325@gmail.com','root','Pere')")

	db.Close()
}


