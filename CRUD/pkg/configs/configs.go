package configs

import(
    "encoding/json"
    "io/ioutil"
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

    jsonFile, loadErr := os.OpenFile("./config/config.json", os.O_CREATE, 0777)

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