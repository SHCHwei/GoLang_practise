package user


type User struct{
    ID int `gorm:"primaryKey" form:PersonId`
    LastName string `form:"LastName"`
    FirstName string `form:"FirstName"`
    Email string `form:"Email"`
    Address string `form:"Address"`
    City string `form:"City"`
}