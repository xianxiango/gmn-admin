// +build !super

package routers

import (
	"gmn/controllers"

	"github.com/astaxie/beego"
)

func init() {
	//admin
	nsAdmin := beego.NewNamespace("/gmn",
		//系统
		beego.NSRouter("/image/upload", &controllers.BaseController{}, "*:UploadImage"),
		beego.NSRouter("/image/delete", &controllers.BaseController{}, "*:DeleteUpload"),

		//广告
		beego.NSRouter("/banner/list", &controllers.BannerController{}, "*:List"),
	)

	beego.AddNamespace(nsAdmin)
}
