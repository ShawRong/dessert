package main

import (
	"fmt"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New() // 实例一个iris对象
	app.RegisterView(iris.HTML("./view", ".html"))
	app.HandleDir("/static", iris.Dir("./view/static"))

	//配置路由
	app.Get("/", func(ctx iris.Context) {
		ctx.View("index.html")
	})

	app.Post("/temp", UserLogin)

	// 路由分组
	party := app.Party("/hello")
	// 此处它的路由地址是: /hello/world
	party.Get("/world", func(ctx iris.Context) {
		ctx.View("index.html")
	})

	// 启动服务器
	app.Run(iris.Addr(":8085"), iris.WithCharset("UTF-8"))
	// 监听地址:本服务器上任意id端口8085,设置字符集utf8
}

type User struct {
	Username string `json:"username"`
	Pwd      string `son:"pwd"`
}

func UserLogin(ctx iris.Context) {
	aul := new(User)
	if err := ctx.ReadJSON(&aul); err != nil {
		ctx.StatusCode(iris.StatusOK)
		fmt.Println("error")
		return
	}

	ctx.StatusCode(iris.StatusOK)
	if aul.Username == "Name" && aul.Pwd == "NAME" {
		ctx.WriteString("Nice")
	} else {
		ctx.WriteString("Bad")
	}
	return
}
