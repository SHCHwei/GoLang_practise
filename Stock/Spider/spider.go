package spider

import (
    "net/http"
    "encoding/json"
    "strings"
    "strconv"
    "time"
    "fmt"
    "log"
    "io"
    "stock/database"
    "go.mongodb.org/mongo-driver/bson"
)



func Run(db *database.MongoDB){

    count, err := db.Collection.CountDocuments(db.Ctx, bson.D{})

    if err != nil{
        logF("db error", err)
    }else if count > 0 {
        logF("Stocks is already exist", nil)
    } else {
        stocksList(db)
    }

    var no string

    fmt.Println("Please input stock id : ")
    fmt.Scan(&no)

    historyData(db, no)


}

func stocksList(db *database.MongoDB) error{

    docs := []interface{}{}

    for _ , no := range apiFilter{

        apiUrl := "https://www.twse.com.tw/rwd/zh/api/codeFilters?filter=" + no + "&_=" + strconv.FormatInt(time.Now().Unix(), 10)
        resp, err := http.Get(apiUrl)

        logF("open api", err)

        defer resp.Body.Close()
        data, err := io.ReadAll(resp.Body)

        logF("open data", err)

        var s stocksAPI

        logF("read json", json.Unmarshal(data, &s))

        for _, v := range s.Result{
            list := strings.Split(v, "\t")
            docs = append(docs, Stocks{Id: list[0], Name: list[1]})
        }

    }

    _, err := db.Collection.InsertMany(db.Ctx, docs)
    logF("data insert db", err)

    return err
}

func historyData(db *database.MongoDB, no string){


    if len([]byte(no)) != 4 && len([]byte(no)) != 5 {
       fmt.Println("input format error")
    }

    fmt.Println("your input stock number : " , no)

    var searchRange bool = false

    for _, v := range years {

        for _, val := range months {

            if v == "2020" && val == "02" {
                searchRange = true
                break
            }

            apiUrl := "https://www.twse.com.tw/rwd/zh/afterTrading/STOCK_DAY?date="+v+val+"01&stockNo="+no+"&response=json&_=" + strconv.FormatInt(time.Now().Unix(), 10)

            resp, err := http.Get(apiUrl)

            logF("open api", err)

            defer resp.Body.Close()
            data, err := io.ReadAll(resp.Body)

            logF("open api data", err)

            var i stockAPI

            json.Unmarshal(data, &i)

            docs := []interface{}{}

            for _ , v := range i.Data{

                strDate := strings.Split(v[0], "/")

                CEYear, _ := strconv.Atoi(strDate[0])
                strDate[0] = strconv.Itoa(1911 + CEYear)

                newDate := strings.Join(strDate, "-")

                docs = append(docs, stockDetail{newDate,v[1],v[2],v[3],v[4],v[5],v[6],v[7],v[8]}  )
            }

            db.Collection = db.Client.Database("stock").Collection("no"+no)

            fmt.Println("your input stock number : " , docs)


            _, errB := db.Collection.InsertMany(db.Ctx, docs)
            logF("data insert db", errB)

        }

        time.Sleep(100 * time.Millisecond)

        if searchRange {
            break
        }

    }

}


func logF(s string , err error){
    if err != nil {
		log.Fatal(s, err.Error())
	}else{
        log.Println(s)
    }
}