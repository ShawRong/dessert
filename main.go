package main

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/kataras/iris/v12"
)

type Response struct {
	Status bool        `json:"status"`
	Data   interface{} `json:"data"`
}

func main() {
	app := iris.New() // 实例一个iris对象
	app.Favicon("./favicons/favicon.ico")
	app.RegisterView(iris.HTML("./view", ".html"))      //资源库
	app.HandleDir("/static", iris.Dir("./view/static")) //静态文件

	//配置路由
	app.Get("/", func(ctx iris.Context) {
		ctx.View("index.html")
	})

	//用户登录
	app.Post("/login", UserLogin)

	//用户登录成功请求下一个界面
	// /user/{name}/home
	userRouter := app.Party("/user")
	userRouter.Get("/{name:string}/home/user.html", func(ctx iris.Context) {
		ctx.ServeFile("./view/user.html")
	})
	userRouter.Get("/{name:string}/home/static/js/vue.js", func(ctx iris.Context) {
		ctx.ServeFile("./view/static/js/vue.js")
	})

	// /user/{name}/getname
	userRouter.Get("/{name:string}/getname", func(ctx iris.Context) {
		name := ctx.Params().Get("name")
		_, _ = ctx.JSON(Response{Status: true, Data: name})
	})

	// /user/{name}/getquestion
	userRouter.Post("/{name:string}/getquestion", RandGenerate)

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

//login
type User struct {
	Username string `json:"usrname"`
	Pwd      string `json:"pwd"`
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
		data := "Nice"
		_, _ = ctx.JSON(Response{Status: true, Data: data})
	} else {
		data := "Bad"
		_, _ = ctx.JSON(Response{Status: true, Data: data})
	}
	return
}

type Question struct {
	X   string `json:"x"`
	Y   string `json:"y"`
	Sig string `json:"sig"`
	Res string `json:"res"`
}

func RandGenerate(ctx iris.Context) {
	x := rand.Intn(100)
	y := rand.Intn(100)
	sig := rand.Intn(2) //0 for + 1 for -
	res := rand.Intn(2)
	var mark string
	if sig == 0 {
		res = x + y
		mark = "+"
	} else {
		res = x - y
		mark = "-"
	}
	_, _ = ctx.JSON(Question{X: strconv.Itoa(x), Y: strconv.Itoa(y), Sig: mark, Res: strconv.Itoa(res)})
}
