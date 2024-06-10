package main

import (
    "github.com/gocolly/colly/v2"
    "fmt"
)


func main(){

    c := colly.NewCollector()

    c.OnResponse(func(r *colly.Response) {
        r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36")
    })

    c.OnHTML("p", func(e *colly.HTMLElement){
        fmt.Println(e.Text)
    })


    c.Visit("https://zerotomastery.io/blog/golang-practice-projects/")
}


