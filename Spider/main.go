package main

import (
    "github.com/gocolly/colly/v2"
    "fmt"
    "time"
    "os"
)

type List struct{
    Title string `json:title`
    Quantity string `json:quantity`
}

func main(){

    var stockID string = ""

    fmt.Println("請輸入股票號碼 : ")
    fmt.Scan(&stockID)

    if stockID != ""{
        runSpider(stockID)
    } else {
        fmt.Println("Fuck you : ")
    }

}


func runSpider(stockID string){

    c := colly.NewCollector()

    c.OnResponse(func(r *colly.Response) {
        r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36")
    })


    c.Limit(&colly.LimitRule{
        DomainGlob:  "*",
        RandomDelay: 2 * time.Second,
    })

    var formData = map[string]string{
        "encodeURIComponent": "1",
        "step": "1",
        "firstin": "1",
        "off": "1",
        "keyword4": "",
        "code": "",
        "TYPEK2": "",
        "checkbtn": "",
        "queryName": "co_id",
        "inpuType": "co_id",
        "TYPEK": "all",
        "isnew": "true",
        "co_id": stockID,
        "year": "",
        "month": "",
    }


    var targetData []List

    c.OnHTML(".hasBorder", func(e *colly.HTMLElement){
        e.ForEach("tr", func(i int, el *colly.HTMLElement){
            if i > 0 {
                targetData = append(targetData, List{el.ChildText("th"), el.ChildText("td")})
            }
        })
    })

    c.OnError(func(r *colly.Response, err error){
        if err != nil {
            fmt.Println()
            fmt.Printf("[Error INFO][%d]%s", r.StatusCode, err)
        }
    })


    err := c.Post("https://mops.twse.com.tw/mops/web/ajax_t05st10_ifrs", formData)

    if err != nil {
        fmt.Println()
        fmt.Printf("[Post error]%s", err)
    }

    fmt.Println("spider success", stockID)
    saveFile(targetData, stockID)
}



func check(e error) {
    if e != nil {
        panic(e)
    }
}

func saveFile(data []List, stockID string){

    f, err := os.Create(stockID+".txt")
    check(err)

    defer f.Close()

    for _, v := range data{
        _, err := f.WriteString(v.Title+":"+v.Quantity+"\n")
        check(err)
    }
}