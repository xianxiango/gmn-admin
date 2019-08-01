package controllers

import (

	// time "gmn-admin/common"
	// "gmn-admin/common/permit"

	time "gmn-admin/common"
	"strconv"
	"strings"

	"github.com/astaxie/beego"

	"gmn-admin/libs"

	// "gmn-admin/common/enum"

	model "gmn-admin/models"
	//"github.com/dchest/captcha"
)

var publicUrl = "http://localhost:8080"

type AdminController struct {
	BaseController
}

//
func (this *AdminController) AjaxLogin() {
	beego.ReadFromRequest(&this.Controller)
	if !this.isPost() {
		this.jsonResult(map[string]interface{}{
			"message": "非POST方法",
			"code":    1,
		})
		return
	}

	username := strings.TrimSpace(this.GetString("username"))
	password := strings.TrimSpace(this.GetString("password"))

	remember, _ := this.GetInt("remember")
	// isagent, _ := this.GetInt("isagent")

	if username == "" || password == "" {
		this.jsonResult(map[string]interface{}{
			"message": "缺少参数",
			"code":    1,
			"data":    map[string]interface{}{},
		})
	}

	user := model.GetAdminByUsername(username)
	errorMsg := ""
	// log.Println("kaka", user.Platform)
	// if isagent == 1 {
	// 	errorMsg = "帐号或密码错误"
	// } else {
	if user.Password != libs.Md5([]byte(password+"cgt123")) {
		errorMsg = "帐号或密码错误"
	} else {
		user.LastLoginIp = this.getClientIp()
		user.Uptime = time.Now()
		user.Update("uptime,last_login_ip")

		authkey := libs.Md5([]byte(this.getClientIp() + "|" + user.Password + "cgt123"))
		if remember == 1 {
			this.Ctx.SetCookie("gmnauth", strconv.Itoa(user.Uid)+"|"+authkey, 7*86400) //7*86400
		} else {
			this.Ctx.SetCookie("gmnauth", strconv.Itoa(user.Uid)+"|"+authkey, 86400) //86400
		}
		this.jsonResult(map[string]interface{}{
			"message": "操作成功",
			"code":    0,
		})
		return
	}
	// }

	this.jsonResult(map[string]interface{}{
		"message": errorMsg,
		"code":    1,
		"data":    map[string]interface{}{},
	})
	return

}

// 退出登录
func (this *AdminController) Logout() {
	this.Ctx.SetCookie("gmnauth", "")
	this.jsonResult(map[string]interface{}{
		"message": "操作成功",
		"code":    0,
	})
}

// func (this *AdminController) PermitGroups() {
// 	pgs, _ := model.FindPermitGroups()
// 	this.jsonSuccResult(pgs)
// }

// func (this *AdminController) NewPermitGroup() {
// 	parentid := this.GetString("parentid")
// 	name := this.GetString("name")
// 	desc := this.GetString("desc")

// 	pg := model.NewPermitGroup(parentid, name, desc)
// 	pg.Insert()
// 	this.jsonSuccResult(pg)
// }

// func (this *AdminController) GetAdmins() {
// 	admins := model.GetAdmins()
// 	ret := []map[string]interface{}{}
// 	for _, v := range admins {
// 		permitGroup := model.GetPermitGroupByGroupId(v.PermitGroupid)
// 		m := ext.Object2Map(v)
// 		m["group_name"] = permitGroup.Name
// 		ret = append(ret, m)
// 	}
// 	this.jsonSuccResult(ret)
// }

// func (this *AdminController) GenerateAdmin() {
// 	username := this.GetString("username")
// 	admin := model.GetAdminByUsername(username)
// 	if admin.Username != "" {
// 		this.jsonFailResult("用户名已存在")
// 		return
// 	}

// 	pwd := common.BuilderPrimaryid()[:8]
// 	newAdmin := model.NewAdmin(username, pwd, "salt")
// 	newAdmin.Password = libs.Md5([]byte(pwd + newAdmin.Salt))
// 	newAdmin.Insert()

// 	adminQuery := model.GetAdminByUsername(username)
// 	//记录日志
// 	adminLog := model.NewAdminLogs(this.adminId)
// 	adminLog.Action = enum.GenerateAdmin
// 	adminLog.Module = "admin"
// 	adminLog.Target = fmt.Sprintf("%v", newAdmin.Uid)
// 	adminLog.Value = fmt.Sprintf("%v", newAdmin.PermitGroupid)
// 	adminLog.Insert()
// 	//返回的密码重新处理
// 	adminQuery.Password = pwd
// 	this.jsonSuccResult(adminQuery)
// }

// func (this *AdminController) RemoveAdmin() {
// 	id, _ := this.GetInt("id")
// 	admin := model.GetAdminById(id)
// 	if admin.Uid == 0 {
// 		this.jsonFailResult("该管理员不存在")
// 		return
// 	}
// 	admin.PermitGroupid = ""
// 	admin.State = 1
// 	admin.Update("state,permit_groupid")
// 	this.jsonSuccResult("")
// }

// func (this *AdminController) ChangePassword() {
// 	id, _ := this.GetInt("id")
// 	oldPsw := this.GetString("old_psw")
// 	newPsw := this.GetString("new_psw")

// 	admin := model.GetAdminById(id)
// 	if admin.Uid == 0 {
// 		this.jsonFailResult("该管理员不存在")
// 		return
// 	}

// 	myPermitGroup := model.GetPermitGroupByGroupId(this.admin.PermitGroupid)
// 	permitGroup := model.GetPermitGroupByGroupId(admin.PermitGroupid)

// 	refresh := func(mgr bool) {
// 		if admin.Password == libs.Md5([]byte(oldPsw+admin.Salt)) || mgr {
// 			admin.Password = libs.Md5([]byte(newPsw + admin.Salt))
// 			admin.Update("password")
// 			this.jsonSuccResult("")
// 		}
// 		adminLog := model.NewAdminLogs(this.adminId)
// 		adminLog.Action = enum.ChangePassword
// 		adminLog.Module = "admin"
// 		adminLog.Target = fmt.Sprintf("%d", admin.Uid)
// 		adminLog.Value = fmt.Sprintf("%s", admin.PermitGroupid)
// 		adminLog.Insert()
// 	}

// 	if this.adminId == id {
// 		refresh(false)
// 	} else if this.admin.Issuper == 1 {
// 		refresh(true)
// 	} else if myPermitGroup.Groupid == permitGroup.Parentid {
// 		refresh(true)
// 	} else {
// 		this.jsonFailResult("权限不够")
// 		return
// 	}
// 	this.jsonSuccResult("操作成功")
// }

// func (this *AdminController) SetAdminPermitGroup() {
// 	uid, _ := this.GetInt("uid")
// 	groupid := this.GetString("groupid")

// 	admin := model.GetAdminById(uid)
// 	if admin.Username == "" {
// 		this.jsonFailResult("找不到该用户")
// 		return
// 	}

// 	adminLog := model.NewAdminLogs(this.adminId)
// 	adminLog.Action = enum.SetAdminPermitGroup
// 	adminLog.Module = "admin"
// 	adminLog.Target = fmt.Sprintf("%d", admin.Uid)
// 	adminLog.Value = fmt.Sprintf("%s", admin.PermitGroupid)
// 	adminLog.Insert()

// 	admin.PermitGroupid = groupid
// 	admin.Uptime = time.Now()
// 	admin.Update("permit_groupid")
// 	this.jsonSuccResult(admin)
// }

// func (this *AdminController) RemoveAdminPermitGroup() {
// 	uid, _ := this.GetInt("uid")
// 	//groupid := this.GetString("groupid")

// 	admin := model.GetAdminById(uid)
// 	if admin.Username == "" {
// 		this.jsonFailResult("找不到该用户")
// 		return
// 	}

// 	adminLog := model.NewAdminLogs(this.adminId)
// 	adminLog.Action = enum.RemoveAdminPermitGroup
// 	adminLog.Module = "admin"
// 	adminLog.Target = fmt.Sprintf("%d", admin.Uid)
// 	adminLog.Value = fmt.Sprintf("%s", admin.PermitGroupid)
// 	adminLog.Insert()

// 	admin.PermitGroupid = ""
// 	admin.Uptime = time.Now()
// 	admin.Update("permit_groupid")
// 	this.jsonSuccResult(admin)
// }

// func (this *AdminController) PermitModules() {
// 	c := &AdminAccessController{}
// 	c.Ctx = this.Ctx
// 	c.PermitModules()
// }

// func (this *AdminController) GetAllModules() {
// 	moduless := model.GetModules()
// 	var items = []map[string]interface{}{}
// 	for _, m := range moduless {
// 		key := m.ModuleName
// 		val := permit.PermitsAction[key]
// 		item := ext.Object2Map(m)
// 		item["actions"] = val
// 		items = append(items, item)
// 	}
// 	this.jsonSuccResult(items)
// }

// //GetModules

// func (this *AdminController) EditPermitMudule() {
// 	groupId := this.GetString("groupid")
// 	module := this.GetString("module")
// 	actions := this.GetString("actions")

// 	pm := model.GetPermitModuleByGroupId(groupId, module)
// 	if pm.Groupid != "" {
// 		if actions == "" {
// 			this.jsonFailResult("模块权限已存在")
// 			return
// 		}

// 		//更新action
// 		pm.Action = actions
// 		pm.Update("action")

// 		//记录操作日志
// 		adminLog := model.NewAdminLogs(this.adminId)
// 		adminLog.Action = enum.EditPermitMudule
// 		adminLog.Module = "admin"
// 		adminLog.Target = pm.Groupid
// 		adminLog.Insert()

// 		this.jsonSuccResult(pm)
// 		return
// 	}

// 	actionString := permit.GetActionsStringByModuleName(module)
// 	pmNew := model.NewPermitModule(groupId, module, actionString)
// 	pmNew.Insert()

// 	adminLog := model.NewAdminLogs(this.adminId)
// 	adminLog.Action = enum.EditPermitMudule
// 	adminLog.Module = "admin"
// 	adminLog.Target = pmNew.Groupid
// 	adminLog.Insert()

// 	this.jsonSuccResult(pmNew)
// }

// func (this *AdminController) RemovePermitMudule() {
// 	groupId := this.GetString("groupid")
// 	topModule := this.GetString("top_module")
// 	module := this.GetString("module")

// 	pm := model.GetPermitModuleByGroupId(groupId, module)
// 	if pm.Groupid == "" {
// 		this.jsonFailResult("找不到该模块")
// 		return
// 	}

// 	var moduleName = module
// 	if pm.Module == "*" {
// 		modules := model.GetModules()
// 		for _, module := range modules {
// 			if module.ModuleName == moduleName {
// 				continue
// 			}

// 			//获取模块的所有Action
// 			var arrayActions = []string{}
// 			actions, _ := permit.PermitsAction[topModule][module.ModuleName]
// 			for _, act := range actions {
// 				arrayActions = append(arrayActions, act.Action)
// 			}

// 			actionsDot := strings.Join(arrayActions, ",")
// 			pmNew := model.NewPermitModule(pm.Groupid, module.ModuleName, actionsDot)
// 			pmNew.Insert()
// 		}
// 	}

// 	adminLog := model.NewAdminLogs(this.adminId)
// 	adminLog.Action = enum.RemovePermitMudule
// 	adminLog.Module = "admin"
// 	adminLog.Target = pm.Groupid
// 	adminLog.Insert()

// 	pm.Delete()
// 	this.jsonSuccResult("")
// }
