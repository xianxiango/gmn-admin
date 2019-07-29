package config

import (
	"fmt"
	center "gmn-admin/db/center"
	_ "gmn-admin/routers"

	"github.com/astaxie/beego"
	//aes "gmn-admin/common/crypto"
)

func LoadConfigDebug() {

	//初始化后台配置
	{
		host := beego.AppConfig.String("adminHost")
		port := beego.AppConfig.String("adminPort")
		usr := beego.AppConfig.String("adminUser")
		psw := beego.AppConfig.String("adminPass")
		//psw = aes.DecryptAes(psw)
		adminDB := beego.AppConfig.String("adminDB")
		addr := fmt.Sprintf("%s:%s", host, port)
		center.InitDatabase(addr, usr, psw, adminDB)
		center.Debug()
	}

}
