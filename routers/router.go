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
		beego.NSRouter("/advert/add", &controllers.AdvertController{}, "*:Add"),
		beego.NSRouter("/advert/list", &controllers.AdvertController{}, "*:List"),
	)

	beego.AddNamespace(nsAdmin)
}
