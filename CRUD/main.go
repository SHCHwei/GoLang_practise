package main


import(
    "crud/router"
    _ "crud/database"

)



func main(){
    router := router.SetupRouter()
    router.Run(":8004")
}


