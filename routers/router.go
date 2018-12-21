package routers

import (
	"github.com/astaxie/beego"
	"AppCodeServer/controllers"
)

func init() {
	//beego.Router("/download/plugin/pluginLogin", &controllers.ApkController{}, "*:DownloadFile")

/*

	var FilterUser = func(ctx *context.Context) {
		//sessionId := ctx.Input.Cookie("sessionId")
		beego.Debug("requestUrl:", ctx.Request.RequestURI)
	}

	beego.InsertFilter("/*", beego.BeforeExec, FilterUser)*/



	beego.Router("/download/plugin/:pluginName([\\w]+)", &controllers.ApkController{},"*:DownloadFile")
	beego.Router("/upload", &controllers.ApkController{}, "get:UpLoadForm")
	beego.Router("/upload", &controllers.ApkController{}, "post:UploadFile")


	clientNs := beego.NewNamespace("/v1",
		/*App版本检测*/
		beego.NSNamespace("/posts",
			beego.NSRouter("/", &controllers.PostController{}, "get:GetPosts"),
		),
	)

	beego.AddNamespace(clientNs)
}
