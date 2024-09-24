package spider

var years = [5]string{"2020", "2021", "2022", "2023", "2024"}
var months = [12]string{"01","02","03","04","05","06","07","08","09","10","11","12"}

var apiFilter = []string{"01","02","03","04","05","06","07","08","09","10","11","12","13","14","15","16","17","18","19","20"}

type stocksAPI struct{
    Filter string `json:filter`
    Result []string `json:result`
}


type Stocks struct{
    Id string `bson:"id"`
    Name string `bson:"name"`
}


type stockAPI struct{
    Data [][]string `json:data`
    Date string `json:date`
    Fields interface{} `json:fields`
    Note interface{} `json:note`
    stat string `json:stat`
    title string `json:title`
    total string `json:total`
}


type stockDetail struct{
    Date string
    TradeVolume string
    TradeValue  string
    OpeningPrice string
    HighestPrice string
    LowestPrice string
    ClosingPrice string
    Change string
    Transaction string
}
