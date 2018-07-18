package routers

import (
	"新建文件夹/novalhtml/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
