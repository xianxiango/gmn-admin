package model

import (
	"fmt"
	db "gmn-admin/db/center"
	"log"
	"time"
)

type Advert struct {
	ID         int       `json:"ID" xorm:"ID"`
	Title      string    `json:"Title" xorm:"Title"`
	Content    string    `json:"Content" xorm:"Content"`
	Url        string    `json:"Url" xorm:"Url"`
	IsShow     int       `json:"IsShow" xorm:"IsShow"`
	Module     int       `json:"Module" xorm:"Module"`
	Sort       int       `json:"Sort" xorm:"Sort"`
	Createtime time.Time `json:"CreateTime" xorm:"CreateTime"`
	Updatetime time.Time `json:"UpdateTime" xorm:"UpdateTime"`
}

func NewImagelist() *Advert {
	return &Advert{}
}

func GetImagelist(Module int, pageSize, offset int) ([]Advert, int64) {
	var ad []Advert
	type T struct {
		Total int64
	}
	var total T
	sql := fmt.Sprintf("select * from imagelist where Module = %d order by CreateTime desc limit %d,%d", Module, offset, pageSize)
	err := db.GetDB().Sql(sql).Find(&ad)
	_, err = db.GetDB().Sql(`select COUNT(*) as Total from imagelist where Module = ? `, Module).Get(&total)
	if err != nil {
		panic(err)
	}
	return ad, total.Total
}

func (c *Advert) Insert() error {
	log.Println(c)
	sql := fmt.Sprintf(`INSERT INTO imagelist (Title, Content,Url,IsShow,Module,Sort) VALUES ('%s','%s','%s',%d,%d,%d)`, c.Title, c.Content, c.Url, c.IsShow, c.Module, c.Sort)
	_, err := db.GetDB().Exec(sql)
	if err != nil {
		panic(err)
	}
	return err
}

func (c *Advert) Update(col string) error {
	_, err := db.GetDB().Where("ID = ?", c.ID).Cols(col).Update(c)

	return err
}

func (c *Advert) Delete() error {
	_, err := db.GetDB().Where("ID = ?", c.ID).Delete(c)
	if err != nil {
		panic(err)
	}
	return err
}
