package controllers

import (
	"contentAPI/models"
	"contentAPI/services"
	"net/http"
    "github.com/gin-gonic/gin"
    "time"
    "log"
)


type ContactController struct {
	CS  services.ContactService
}

type Search struct {
    StartDate time.Time `form:"startDate"`
    EndDate time.Time `form:"endDate"`
}

// 初始化"控制"用struct  放入"服務"
func New(cservice services.ContactService) ContactController {

	return ContactController{
		CS: cservice,
	}
}


//路徑註冊
func (u *ContactController) RegisterContactRoutes(r *gin.RouterGroup) {
	userroute := r.Group("/user")
	userroute.POST("/create", u.CreateContact)
	userroute.GET("/search", u.SearchContact)
}


func(u *ContactController) CreateContact(ctx *gin.Context){

    var formData models.ContactOne

    ctx.Bind(&formData)

    now := time.Now()
    formData.Time = now.UTC()
    result := u.CS.CreateContact(&formData)

    if result == nil {
        ctx.JSON(http.StatusOK, gin.H{"message": result})
    } else {
        ctx.JSON(http.StatusOK, gin.H{"message": "sorry~ 傳送失敗"})
    }
}


func(u *ContactController) SearchContact(ctx *gin.Context){

    var timeLimit Search
    var defaultTime time.Time

    ctx.Bind(&timeLimit)

    now := time.Now()
    year, month, day := now.Date()

    if defaultTime.Equal(timeLimit.StartDate) {
        timeLimit.StartDate = time.Date(year, month, day, 00, 00, 00, 0, time.UTC)
    }

    if defaultTime.Equal(timeLimit.EndDate) {
        timeLimit.EndDate = time.Date(year, month, day, 23, 59, 59, 0, time.UTC)
    }

    log.Println(timeLimit)

    var data []models.Contacts
    var err error
    data, err = u.CS.SearchContact(timeLimit.StartDate, timeLimit.EndDate);

    ctx.JSON(http.StatusOK, gin.H{"message": err, "data": data})

}

