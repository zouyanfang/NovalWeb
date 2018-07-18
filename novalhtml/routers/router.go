package routers

import (
	"NovalWeb/novalhtml/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
