// +build !super

package routers

import (
	"gmn-admin/controllers"

	"github.com/astaxie/beego"
)

func init() {
	//admin
	nsAdmin := beego.NewNamespace("/gmn-admin",
		//系统
		beego.NSRouter("/image/upload", &controllers.BaseController{}, "*:UploadImage"),
		beego.NSRouter("/image/delete", &controllers.BaseController{}, "*:DeleteUpload"),

		//广告
		beego.NSRouter("/banner/list", &controllers.BannerController{}, "*:List"),
	)

	beego.AddNamespace(nsAdmin)
}
