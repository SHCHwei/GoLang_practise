package main

import(
    "stock/database"
    "stock/spider"
)

var Mongo database.MongoDB


func main(){

    database.Con(&Mongo)

    defer database.DisCon(&Mongo)

    spider.Run(&Mongo)

}


