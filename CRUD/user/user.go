package user

import(
    mf "crud/messageFormat"
    "github.com/gin-gonic/gin"
    //"github.com/gin-gonic/gin/binding"
    //"github.com/go-playground/validator/v10"
    db "crud/database"
    "encoding/json"
    "fmt"
)



type InputUser struct{
    LastName string `form:"LastName" binding:"required"`
    FirstName string `form:"FirstName" binding:"required"`
    Email string `form:"Email" binding:"required"`
    Address string `form:"Address" binding:"required"`
    City string `form:"City" binding:"required"`
    Token string `form:"Token" binding:"required"`
}

type User struct{
    PersonId int
    LastName string
    FirstName string
    Email string
    Address string
    City string
}


func Create(c *gin.Context){

    var someone InputUser

    c.ShouldBind(&someone)

    new_guy := User{FirstName: someone.FirstName, LastName: someone.LastName, Email: someone.Email, Address: someone.Address, City: someone.City,}

    result := db.MariaDB.Table("users").Create(&new_guy)

    if result.Error != nil {
        fmt.Println("show the error message => ", result.Error)
        c.JSON(200, mf.MsgFormat.DBFailed())
    }else{
        c.JSON(200, mf.MsgFormat.Success(""))
    }
}

func Read(c *gin.Context){
    var user User

    PersonId := c.PostForm("PersonId")

    result := db.MariaDB.First(&user, "person_id = ?", PersonId)

    if result.Error != nil {
        c.JSON(200, mf.MsgFormat.DBFailed())
    }else{

        jsonResp, _ := json.Marshal(user)
        c.JSON(200, mf.MsgFormat.Success(string(jsonResp)))
    }
}


func Update(c *gin.Context){

//     var someone InputUser
//
//     c.ShouldBind(&someone)
//
//     guy := User{FirstName: someone.FirstName, LastName: someone.LastName, Email: someone.Email, Address: someone.Address, City: someone.City,}
//
//     db.Model(&user).Where("person_id = ?", PersonId).Updates(guy)
//
//     if result.Error != nil {
//
//         c.JSON(200, mf.MsgFormat.DBFailed())
//     }else{
        c.JSON(200, mf.MsgFormat.Success(""))
//     }
}

func Delete(c *gin.Context){

    var user User

    PersonId := c.PostForm("PersonId")

    result := db.MariaDB.Delete(&user, PersonId)

    if result.Error != nil {
        c.JSON(200, mf.MsgFormat.DBFailed())
    }else{
        c.JSON(200, mf.MsgFormat.Success(""))
    }
}