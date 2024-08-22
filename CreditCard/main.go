package main

import (
    "log"
    "strings"
    "strconv"
    "fmt"
    "net/http"
    "encoding/json"
)


func main(){

    http.HandleFunc("/vail", vail)

    logF("Server Run...", nil)

    err := http.ListenAndServe(":1002", nil)

    if err != nil {
        log.Fatal("Server Error", err)
    }
}


func vail(w http.ResponseWriter, r *http.Request){

    type cardInfo struct{
        CardNumber string `json:"cardNumber"`
    }

    //r.Body = w.MaxBytesReader(w, r.Body(), 1048576)

    nd := json.NewDecoder(r.Body)
    nd.DisallowUnknownFields()

    var c cardInfo

    err := nd.Decode(&c)

    if err != nil {
        logF("input format is error", err)
    }

    vailResult := LuhnV2(c.CardNumber)

    if vailResult {
        w.Write(buildResponse(map[string]string{"statue": "true", "message": "驗證成功"}))
    }else{
        w.Write(buildResponse(map[string]string{"statue": "false", "message": "驗證失敗"}))
    }

}

func LuhnV2(cardNumber string) bool {

    var inputCheck string = cardNumber[len(cardNumber)-1:len(cardNumber)]
    var totalNum int64

    for i, val := range strings.Split(cardNumber[:len(cardNumber)-1], "") {

        newNum, _ := strconv.ParseInt(val, 10, 64)

         if i != 0 && i % 2 == 1 {

            newNum = newNum * 2

            if newNum > 9{
                a := newNum / 10
                b := newNum % 10
                newNum = a+b
            }
        }

        totalNum = totalNum + newNum
    }

    myCheckNumber := 10 - (totalNum % 10)
    checkNumber, _ := strconv.ParseInt(inputCheck, 10, 64)

    fmt.Println(checkNumber == myCheckNumber)

    return checkNumber == myCheckNumber
}



func LuhnV1(cardNumber string){

    list := strings.Split(cardNumber, "")

    var newList []int64
    var totalNum int64

    for i, val := range list {

        newNum, _ := strconv.ParseInt(val, 10, 64)

         if i != 0 && i % 2 == 1 {

            newNum = newNum * 2

            if newNum > 9{
                a := newNum / 10
                b := newNum % 10
                newNum = a+b

            }
        }
        newList = append(newList, newNum)
        totalNum = totalNum + newNum
    }

    checkNumber := 10 - (totalNum % 10)

    fmt.Printf("newList => %v, totalNum => %d, check => %d", newList, totalNum, checkNumber)

}

func logF(s string , err error){
    if err != nil {
		log.Fatal(s, err.Error())
	}else{
        log.Println(s)
    }
}



func buildResponse(temp map[string]string)([]byte){

    jsonResp, err := json.Marshal(temp)
	if err != nil {
		logF("Error happened in JSON marshal err:", err)
	}

    return jsonResp
}