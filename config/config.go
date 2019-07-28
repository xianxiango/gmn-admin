package config

import (
	_ "gmn-admin/routers"
	//aes "gmn-admin/common/crypto"
)

func LoadConfigDebug() {

	//初始化后台配置
	// {
	// 	host := beego.AppConfig.String("adminHost")
	// 	port := beego.AppConfig.String("adminPort")
	// 	usr := beego.AppConfig.String("adminUser")
	// 	psw := beego.AppConfig.String("adminPass")
	// 	//psw = aes.DecryptAes(psw)
	// 	adminDB := beego.AppConfig.String("adminDB")
	// 	addr := fmt.Sprintf("%s:%s", host, port)
	// 	db.InitDatabase(addr, usr, psw, adminDB)
	// 	db.Debug()
	// }

}
