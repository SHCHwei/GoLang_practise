package user

import(
    mf "crud/messageFormat"
    "github.com/gin-gonic/gin"
    db "crud/database"
    "encoding/json"
    "strconv"
    "fmt"
)



func Create(c *gin.Context){

    var user User

    c.ShouldBind(&user)

    result := db.MariaDB.Create(&user)

    if result.Error != nil {
        fmt.Println("show the error message => ", result.Error)
        c.JSON(200, mf.MsgFormat.DBFailed())
    }else{
        c.JSON(200, mf.MsgFormat.Success(""))
    }
}

func Read(c *gin.Context){

    var user User

    user.ID, _ = strconv.Atoi(c.PostForm("PersonId"))

    result := db.MariaDB.First(&user)

    if result.Error != nil {
        c.JSON(200, mf.MsgFormat.DBFailed())
    }else{
        jsonResp, _ := json.Marshal(user)
        c.JSON(200, mf.MsgFormat.Success(string(jsonResp)))
    }
}


func Update(c *gin.Context){

    var user User

    c.ShouldBind(&user)

    result := db.MariaDB.Save(&user)

    if result.Error != nil {
        fmt.Println(" error => ", result.Error)
        c.JSON(200, mf.MsgFormat.DBFailed())
    }else{
        c.JSON(200, mf.MsgFormat.Success(""))
    }
}

func Delete(c *gin.Context){

    var user User

    user.ID, _ = strconv.Atoi(c.PostForm("PersonId"))

    result := db.MariaDB.Delete(&user)

    if result.Error != nil {
        c.JSON(200, mf.MsgFormat.DBFailed())
    }else{
        c.JSON(200, mf.MsgFormat.Success(""))
    }
}