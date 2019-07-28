package main

import (
	// "flag"
	"gmn-admin/config"
	_ "gmn-admin/routers"

	"github.com/astaxie/beego"
)

// var mode = flag.String("mode", "debug", "debug / release")
// var etcd = flag.String("etcdAddr", "127.0.0.1:11112", "etcd address")

func main() {
	// flag.Parse()
	config.LoadConfigDebug()

	beego.SetStaticPath("/admin", "./static")

	beego.Run()
}
