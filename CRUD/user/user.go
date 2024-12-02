package user

import(
    Rf "crud/pkg/responseFormat"
    "github.com/gin-gonic/gin"
    db "crud/database"
    "encoding/json"
)


func Create(c *gin.Context){

    var user User

    c.ShouldBind(&user)

    result := db.MariaDB.Create(&user)

    if result.Error != nil {
        c.JSON(200, Rf.MsgFormat.DBFailed(result.Error))
    }else{
        c.JSON(200, Rf.MsgFormat.Success(result.Error))
    }
}

func Read(c *gin.Context){

    user := User{Uuid: c.PostForm("UUID")}

    result := db.MariaDB.First(&user)

    if result.Error != nil {
        c.JSON(200, Rf.MsgFormat.DBFailed(result.Error))
    }else{
        jsonResp, _ := json.Marshal(user)
        c.JSON(200, Rf.MsgFormat.Success(string(jsonResp)))
    }
}


func Update(c *gin.Context){

    var user User

    c.ShouldBind(&user)

    result := db.MariaDB.Save(&user)

    if result.Error != nil {
        c.JSON(200, Rf.MsgFormat.DBFailed(result.Error))
    }else{
        c.JSON(200, Rf.MsgFormat.Success(result.Error))
    }
}

func Delete(c *gin.Context){

    user := User{Uuid: c.PostForm("UUID")}

    result := db.MariaDB.Delete(&user)

    if result.Error != nil {
        c.JSON(200, Rf.MsgFormat.DBFailed(result.Error))
    }else{
        c.JSON(200, Rf.MsgFormat.Success(result.Error))
    }
}