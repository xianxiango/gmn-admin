package controllers

import (
	center "gmn-admin/models/center"
)

type AdvertController struct {
	BaseController
}

func (this *AdvertController) List() { //获取列表
	module, _ := this.GetInt("module")
	activities, total := center.GetImagelist(module, this.Page.PageSize, this.Page.Offset)

	this.jsonResult(map[string]interface{}{
		"message": "操作成功",
		"code":    0,
		"data": map[string]interface{}{
			"list":      activities,
			"total":     total,
			"page_no":   this.Page.PageNo,
			"page_size": this.Page.PageSize,
		},
	})
}
func (this *AdvertController) Add() { //新增
	// id, _ := this.GetInt("id")
	title := this.GetString("title")
	content := this.GetString("content")
	url := this.GetString("url")
	isShow, _ := this.GetInt("isShow")
	module, _ := this.GetInt("module")

	advert := center.NewImagelist()

	advert.Title = title
	advert.Content = content
	advert.Module = module
	advert.IsShow = isShow
	advert.Url = url
	err := advert.Insert()
	if err != nil {
		panic(err)
	}
	this.jsonResult(map[string]interface{}{
		"message": "操作成功",
		"code":    0,
	})
}

// func (this *AdvertController) Edit() { //编辑
// 	id, _ := this.GetInt("id")
// 	title := this.GetString("title")
// 	lang := this.GetString("lang")
// 	show, _ := this.GetInt("show")
// 	targetapp := this.GetString("targetapp")
// 	targetmobile := this.GetString("targetmobile")
// 	targetpc := this.GetString("targetpc")
// 	sortapp, _ := this.GetInt("sortapp")
// 	sortmobile, _ := this.GetInt("sortmobile")
// 	sortpc, _ := this.GetInt("sortpc")
// 	imagepc := this.GetString("imagepc")
// 	imageweb := this.GetString("imageweb")
// 	imageios := this.GetString("imageios")
// 	imageandroid := this.GetString("imageandroid")

// 	Banner := center.GetBannerById(id)
// 	if Banner.Id == 0 {
// 		this.jsonFailResult("找不到该记录")
// 		return
// 	}

// 	{
// 		Banner.Title = title
// 		Banner.Lang = lang
// 		Banner.Show = show
// 		Banner.Targetapp = targetapp
// 		Banner.Targetmobile = targetmobile
// 		Banner.Targetpc = targetpc
// 		Banner.Sortapp = sortapp
// 		Banner.Sortmobile = sortmobile
// 		Banner.Sortpc = sortpc
// 		Banner.Imageios = imageios
// 		Banner.Imageandroid = imageandroid
// 		Banner.Imageweb = imageweb
// 		Banner.Imagepc = imagepc
// 		Banner.Updatetime = time.Now()
// 		err := Banner.Update("Title,Targetapp,Targetmobile,Targetpc,Lang,Show,SortApp,SortMobile,Sortpc,Imageios,Imageandroid,,Imageweb,Imagepc,Updatetime")
// 		if err != nil {
// 			msg = map[string]interface{}{
// 				"message": err,
// 				"code":    502,
// 			}
// 		} else {
// 			api.AdminBanner()
// 		}
// 		this.jsonResult(msg)
// 	}

// 	this.jsonResult(map[string]interface{}{
// 		"message": "操作成功",
// 		"code":    0,
// 	})
// }

// func (this *BannerController) Delete() { //删除
// 	id, _ := this.GetInt("id")

// 	Banner := center.GetBannerById(id)
// 	if Banner.Id == 0 {
// 		this.jsonFailResult("找不到该记录")
// 		return
// 	}

// 	err := Banner.Delete()
// 	if err != nil {
// 		msg = map[string]interface{}{
// 			"message": err,
// 			"code":    502,
// 		}
// 	} else {
// 		api.AdminBanner()
// 	}
// 	this.jsonResult(msg)
// }
