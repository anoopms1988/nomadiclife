// @APIVersion 1.0.0
// @Title Nomadiclife
// @Description Backend API for nomadiclife application
// @Contact Anoop MS
package routers

import (
	"nomadiclife/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/signup", &controllers.UserController{})
	beego.Router("/signup", &controllers.UserController{})
	beego.Router("/forget_password", &controllers.UserController{}, "*:ForgetPassword")
}
