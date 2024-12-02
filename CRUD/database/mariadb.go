package database

import(
    "crud/pkg/configs"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "log"
)

func init(){

    var err error

    setting := configs.DBConfig.Maria

    dsn := setting.Account + ":" + setting.Pwd + "@tcp(" + setting.Host + ":" + setting.Port + ")/" + setting.DB + "?charset=utf8mb4&parseTime=True&loc=Local"
    MariaDB , err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

    if err != nil {
        log.Printf("DB Failed : %v", err)
    }
}

