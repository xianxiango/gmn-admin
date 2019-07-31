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
		beego.NSRouter("/advert/search", &controllers.AdvertController{}, "*:Search"),
		beego.NSRouter("/advert/edit", &controllers.AdvertController{}, "*:Edit"),
		beego.NSRouter("/advert/isshow", &controllers.AdvertController{}, "*:IsShow"),
		beego.NSRouter("/advert/del", &controllers.AdvertController{}, "*:Del"),
		beego.NSRouter("/advert/settop", &controllers.AdvertController{}, "*:SetTop"),
		beego.NSRouter("/advert/gettext", &controllers.AdvertController{}, "*:GetTextList"),
		beego.NSRouter("/advert/settext", &controllers.AdvertController{}, "*:SetTextList"),
	)

	beego.AddNamespace(nsAdmin)
}
