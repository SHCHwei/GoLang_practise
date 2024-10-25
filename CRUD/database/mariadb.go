package database

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "log"
    "fmt"
)



func init(){

    var err error

    dsn := "root:admin01@tcp(127.0.0.1:3306)/my_data?charset=utf8mb4&parseTime=True&loc=Local"
    MariaDB , err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

    if err != nil {
        log.Printf("Failed : %v", err)
    }

}

func debug(){

    fmt.Println(MariaDB.Debug)
}