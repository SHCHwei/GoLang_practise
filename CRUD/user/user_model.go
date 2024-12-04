package user


type User struct{
    UUID string `gorm:"column:uuid;primaryKey" form:UUID`
    FirstName string `gorm:"column:firstName" form:"FirstName"`
    LastName string `gorm:"column:lastName" form:"LastName"`
    Gender string `gorm:"column:gender" form:"Gender"`
    Email string `gorm:"column:email" form:"Email"`
    Address string `gorm:"column:address" form:"Address"`
    City string `gorm:"column:city" form:"City"`
//     CreatedAt int64 `gorm:"autoCreateTime"`
//     UpdatedAt int64 `gorm:"autoUpdateTime"`
}