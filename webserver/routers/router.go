package routers

import (
	"github.com/adamar/chaosd/webserver/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "get:Get")


    api := beego.NewNamespace("/api/v1",
       beego.NSRouter("/", &controllers.MainController{}, "get:ApiV1Get"),
    )
    beego.AddNamespace(api)

}
