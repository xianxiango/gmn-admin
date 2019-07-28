package controllers

type BannerController struct {
	BaseController
}

func (this *BannerController) List() { //获取列表
	show, _ := this.GetInt("show")
	// activities, total := center.GetBanners(this.Page.Offset, this.Page.PageSize, show)

	this.jsonResult(map[string]interface{}{
		"message": "操作成功",
		"code":    0,
		"data": map[string]interface{}{
			"list":      show,
			"total":     2,
			"page_no":   this.Page.PageNo,
			"page_size": this.Page.PageSize,
		},
	})
}

// func (this *BannerController) Edit() { //编辑
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
