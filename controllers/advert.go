package controllers

import (
	center "gmn-admin/models/center"
)

type AdvertController struct {
	BaseController
}

func (this *AdvertController) List() { //获取列表
	this.Search()
}
func (this *AdvertController) Search() { //获取列表
	module, _ := this.GetInt("module")
	isShow, _ := this.GetInt("isShow")
	activities, total := center.GetImagelist(module, isShow, this.Page.PageSize, this.Page.Offset)

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

func (this *AdvertController) Edit() { //新增
	id, _ := this.GetInt("id")
	title := this.GetString("title")
	content := this.GetString("content")
	url := this.GetString("url")
	isShow, _ := this.GetInt("isShow")
	// module, _ := this.GetInt("module")

	advert := center.NewImagelist()

	advert.Title = title
	advert.Content = content
	advert.IsShow = isShow
	advert.Url = url
	advert.ID = id
	err := advert.Update("Url,Title,Content,IsShow")
	if err != nil {
		panic(err)
	}
	this.jsonResult(map[string]interface{}{
		"message": "操作成功",
		"code":    0,
	})
}

func (this *AdvertController) Del() { //新增
	id, _ := this.GetInt("id")
	// module, _ := this.GetInt("module")

	advert := center.NewImagelist()

	advert.ID = id
	err := advert.Delete()
	if err != nil {
		panic(err)
	}
	this.jsonResult(map[string]interface{}{
		"message": "操作成功",
		"code":    0,
	})
}

func (this *AdvertController) IsShow() { //新增
	id, _ := this.GetInt("id")
	isShow, _ := this.GetInt("isShow")
	// module, _ := this.GetInt("module")

	advert := center.NewImagelist()

	advert.IsShow = isShow
	advert.ID = id
	err := advert.Update("IsShow")
	if err != nil {
		panic(err)
	}
	this.jsonResult(map[string]interface{}{
		"message": "操作成功",
		"code":    0,
	})
}

func (this *AdvertController) SetTop() { //新增
	id, _ := this.GetInt("id")
	module, _ := this.GetInt("module")
	// module, _ := this.GetInt("module")

	maxSort := center.GetFindSort(module)

	advert := center.NewImagelist()

	advert.Sort = maxSort + 1
	advert.ID = id
	err := advert.Update("Sort")
	if err != nil {
		panic(err)
	}
	this.jsonResult(map[string]interface{}{
		"message": "操作成功",
		"code":    0,
	})
}
