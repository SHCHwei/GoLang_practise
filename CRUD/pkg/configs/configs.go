package configs

import(
    "encoding/json"
    "io/ioutil"
    "strings"
    "os"
    "log"
)


type MariaDB struct{
    Host string `json:"host"`
    Port string `json:"port"`
    Account string `json:"account"`
    Pwd string `json:"pwd"`
    DB string `json:"db"`
}

type DatabaseConfig struct{
    Maria MariaDB `json:"mariadb"`
}

var DBConfig DatabaseConfig

func init(){

    var cfgPath string

    // 判別 是執行測試還是正常使用，因為當前目錄會不一樣
    if len(os.Args) > 1 && strings.HasPrefix(os.Args[1], "-test"){
        cfgPath = "../config/config.json"
    }else{
        cfgPath = "./config/config.json"
    }

    jsonFile, loadErr := os.OpenFile(cfgPath, os.O_CREATE, 0777)

    if loadErr != nil {
        log.Printf("Config load Failed : %v", loadErr)
        panic(loadErr)
    }

    defer jsonFile.Close()

    dataList, readErr := ioutil.ReadAll(jsonFile)

    if readErr != nil {
        log.Printf("Config load Failed : %v", readErr)
    }

	json.Unmarshal(dataList, &DBConfig)
}