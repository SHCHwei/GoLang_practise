package dbcon

import (

	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//資料庫撈取資料成功
//但是拆分func失敗於 全域變數 db的型別

var db 


func DBT(){
	
	dbSettings := fmt.Sprintf("%s:%s@(%s)/%s","root","1324564999","192.168.99.137","finrpt")
	db, err := sql.Open("mysql",dbSettings)
		
	if err != nil{
		panic(err)
	} else {
		db.SetMaxOpenConns(1)
		Select("Select passwd from user limit 1")
	}
}


func Select(sqlString string){

	rows, err := db.Query(sqlString)

	if err != nil{
		panic(err)
	}

	for rows.Next() {
		var passwd string
		err = rows.Scan(&passwd)
		if err != nil{
			panic(err)
		}

		fmt.Println(passwd)
	}
}