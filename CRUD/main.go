package main


import(
    "github.com/gin-gonic/gin"
    "crud/middleware"
    "crud/user"
    _ "crud/database"
)



func main(){

    router := gin.Default()
    router.Use(middleware.Token())

	r := router.Group("/user")
	{
		r.POST("/create", user.Create)
        r.POST("/update", user.Update)
        r.POST("/delete", user.Delete)
        r.POST("/read", user.Read)
	}

    router.Run(":8004")
}


