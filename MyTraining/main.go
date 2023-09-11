package main

import (
	"fmt"
	"abc.com/flagT"
    "abc.com/dbcon"
)

func init(){
    dbcon.DBT()
}

func main(){

	fmt.Println("Hello world")
    flagT.GoflagT()

    // iris.New() 和 iris.Default()
    // iris.New() -> Creates an iris application without any middleware by default
    //app := iris.New()
    app := iris.Default()

    //傳入index方法
    app.Get("/",index)

    //建立 localhost:8090/user/{name}路徑，並使用匿名方法直接處理
    app.Get("/user/{name}", func(ctx iris.Context) {
         name := ctx.Params().Get("name")
         ctx.Writef("Hello %s", name) //顯示 "hello 某某某"字樣
    })

   //建立 localhost:8090/boss/{name}路徑
    app.Get("/boss/{name}", func(ctx iris.Context) {
        name := ctx.Params().Get("name")
        ctx.Writef("Hello %s", name)
    })

    //建立 localhost:8090/welcome?firstname=Jane&lastname=Doe 路徑
    app.Get("/welcome", func(ctx iris.Context) {
        //取得GET值
        // URLParamDefault ("目標名稱","如果不存在 預設值" )
        firstname := ctx.URLParamDefault("firstname", "g")
        lastname := ctx.URLParam("lastname")

        ctx.Writef("Hello %s %s", firstname, lastname)
    })

    //建立 POST localhost:8090/form_post 路徑
    app.Post("/form_post", func(ctx iris.Context) {
        message := ctx.PostValue("message")
        nick := ctx.PostValueDefault("nick", "anonymous")

        ctx.JSON(iris.Map{
            "nick":    nick,
            "message": message,
            "status":  "posted",

        })
    })

    //群組路徑 localhost:8090/v1/loginApp or /loginApp2
    //並送入對應 func
    v1 := app.Party("/v1")
    {
        v1.Post("/loginApp", loginEndpoint)
        v1.Post("/loginApp2", loginEndpoint)
    }

    // 全域中介層(Global middleware) 使用 UseRouter。全域中介層可能指的是如 recover.New() 這種別人寫好會常用的套件
    // recover.New -> 該方法是屬於中介層，從任何錯誤(any panics)和500錯誤上恢復
    app.UseRouter(recover.New())

    //Get("路徑名稱",中介層方法,執行作業)
    app.Get("/benchmark", MyBenchLogger, benchEndpoint)


    // 這個做法應該比較少用到，先建立路由群組，再共同使用一個中介層(如下AuthRequired() )
    // 之後才到路徑名稱和對應功能
    authorized := app.Party("/")
    authorized.Use(AuthRequired())
    {
        authorized.Post("/login", benchEndpoint)
        authorized.Post("/submit", benchEndpoint)
        authorized.Post("/read", benchEndpoint)

        // nested group
        testing := authorized.Party("testing")
        testing.Get("/analytics", benchEndpoint)
    }


    app.Listen(":8090")

}


func index(ctx iris.Context){
    //回傳json型態
    ctx.JSON(iris.Map{
        "status":  "posted",
    })
}

func loginEndpoint(ctx iris.Context){

    message := ctx.PostValue("message")
    nick := ctx.PostValueDefault("nick", "anonymous")

    ctx.JSON(iris.Map{
        "nick":    nick,
        "message": message,
        "status":  "posted",
    })
}

// 中介層方法
func MyBenchLogger(ctx iris.Context){
    fmt.Println("DemoMiddleware")

    //執行下一個處理方法
    ctx.Next()
}

func benchEndpoint(ctx iris.Context){
    ctx.JSON(iris.Map{
        "func": "benchEndpoint",
    })
}