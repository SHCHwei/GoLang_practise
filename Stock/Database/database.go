package database

type DataBase interface{
    connect()
    disConnect()
}

func Con(db DataBase){
    db.connect()
}

func DisCon(db DataBase){
    db.disConnect()
}

