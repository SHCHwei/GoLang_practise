package database

import(
    "crud/pkg/configs"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "log"
    "fmt"
)

func init(){

    var err error

    setting := configs.DBConfig.Maria

    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", setting.Account, setting.Pwd, setting.Host, setting.Port, setting.DB)
    MariaDB , err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

    if err != nil {
        log.Printf("DB Failed : %v", err)
    }
}

