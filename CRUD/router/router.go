package router


import (
    "github.com/gin-gonic/gin"
    "crud/user"
    "crud/middleware"
)



func SetupRouter() *gin.Engine {

	router := gin.Default()
	router.Use(middleware.Token())

	r := router.Group("/user")
	{
		r.POST("/create", user.Create)
        r.POST("/update", user.Update)
        r.POST("/delete", user.Delete)
        r.POST("/read", user.Read)
	}

	return router
}
