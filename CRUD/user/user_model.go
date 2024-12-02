package user


type User struct{
    Uuid string `gorm:"primaryKey" form:ID`
    FirstName string `form:"FirstName"`
    LastName string `form:"LastName"`
    Gender string `form:"Gender"`
    Email string `form:"Email"`
    Address string `form:"Address"`
    City string `form:"City"`
}