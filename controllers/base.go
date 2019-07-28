package controllers

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	systime "time"

	"github.com/astaxie/beego"
	"github.com/tealeg/xlsx"
)

const (
	MSG_OK  = 0
	MSG_ERR = -1
)

type BaseController struct {
	beego.Controller
	controllerName string
	actionName     string
	adminId        int
	userName       string

	Uid int64

	Page struct {
		PageSize int
		PageNo   int
		Offset   int
	}
}

// 重定向
func (this *BaseController) redirect(url string) {
	this.Redirect(url, 302)
	this.StopRun()
}

// 是否POST提交
func (this *BaseController) isPost() bool {
	return this.Ctx.Request.Method == "POST"
}

// 输出json
func (this *BaseController) jsonResult(out interface{}) {
	this.Data["json"] = out
	this.ServeJSON()
	this.StopRun()
}

// 输出json
func (this *BaseController) JsonResult(out interface{}) {
	this.Data["json"] = out
	this.ServeJSON()
	this.StopRun()
}

// 输出json
func (this *BaseController) jsonsResult(out interface{}) {
	this.Data["json"] = out
	this.ServeJSON()
}

// 输出json
func (this *BaseController) jsonSuccResult(data interface{}) {
	this.Data["json"] = map[string]interface{}{
		"code":    0,
		"data":    data,
		"message": "操作成功",
	}
	this.ServeJSON()
	this.StopRun()
}

// 输出json
func (this *BaseController) jsonFailResult(errorMsg string) {
	this.Data["json"] = map[string]interface{}{
		"code":    1,
		"data":    "",
		"message": errorMsg,
	}
	this.ServeJSON()
	this.StopRun()
}

// 输出json
func (this *BaseController) RespJson(out interface{}) {
	this.Data["json"] = out
	this.ServeJSON()
	this.StopRun()
}

func (this *BaseController) UploadImage() {
	log.Print("UploadImage Enter")
	this.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
	this.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Methods", "POST,GET,OPTIONS,DELETE")              //支持的http 动作
	this.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Headers", "x-requested-with,multipart/form-data") //响应头 请按照自己需求添加。

	fd, header, err := this.Ctx.Request.FormFile("file")

	if err != nil {
		this.jsonResult(err.Error())
		return
	}

	ext := filepath.Ext(header.Filename)

	fileId := systime.Now().Unix()
	picId := fmt.Sprintf("%d", fileId)

	prefix := beego.AppConfig.String("filePrefix")
	p := "./static/upload/"
	filename := p + picId + ext

	export := prefix + picId + ext

	os.MkdirAll(p, os.ModeAppend)
	osfd, err := os.Create(filename)
	if err != nil {
		this.jsonResult(err.Error())
		return
	}
	io.Copy(osfd, fd)

	fd.Close()
	osfd.Close()

	this.jsonSuccResult(map[string]interface{}{
		"filename": export,
		"id":       fileId,
	})
}

func (this *BaseController) DeleteUpload() {
	ext := this.GetString("id")
	p := "./static/upload/"

	filename := p + ext

	err := os.Remove(filename)
	if err != nil {
		this.jsonFailResult(err.Error())
		return
	}
	this.jsonSuccResult("操作成功")
}

func (this *BaseController) ExportXlsx(data interface{}) {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	var err error
	file = xlsx.NewFile()
	sheet, err = file.AddSheet("Sheet1")
	if err != nil {
		fmt.Printf("www", err.Error())
	}

	for _, v := range data.([]map[string]interface{}) {
		row = sheet.AddRow()

		var value string
		for _, vt := range v {
			cell = row.AddCell()
			switch vt.(type) {

			case string:
				value = vt.(string)
			case float64:
				value = strconv.FormatFloat(vt.(float64), 'E', -1, 64)
			case int64:
				value = strconv.FormatInt(vt.(int64), 10)
			default:

			}
			cell.Value = value
		}
	}
	err = file.Save("MyXLSXFile.xlsx")
	if err != nil {
		fmt.Printf("eee", err.Error())
	}

}
