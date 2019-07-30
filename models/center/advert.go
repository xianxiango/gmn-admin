package model

import (
	"fmt"
	db "gmn-admin/db/center"
	"log"
	"time"
)

type Imagelist struct {
	ID         int       `json:"ID" xorm:"ID"`
	Title      string    `json:"Title" xorm:"Title"`
	Content    string    `json:"Content" xorm:"Content"`
	Url        string    `json:"Url" xorm:"Url"`
	IsShow     int       `json:"IsShow" xorm:"IsShow"`
	Module     int       `json:"Module" xorm:"Module"`
	Sort       int64     `json:"Sort" xorm:"Sort"`
	Createtime time.Time `json:"CreateTime" xorm:"CreateTime"`
	Updatetime time.Time `json:"UpdateTime" xorm:"UpdateTime"`
}

func NewImagelist() *Imagelist {
	return &Imagelist{}
}

func GetImagelist(Module, IsShow int, pageSize, offset int) ([]Imagelist, int64) {
	var ad []Imagelist
	type T struct {
		Total int64
	}
	var total T
	where := "1=1"
	if Module != 0 {
		where += fmt.Sprintf(" AND Module = %d ", Module)
	}
	if IsShow != -1 {
		where += fmt.Sprintf(" AND IsShow = %d ", IsShow)
	}
	sql := fmt.Sprintf("select * from imagelist where %s order by Sort desc,CreateTime desc limit %d,%d", where, offset, pageSize)
	totalsql := fmt.Sprintf(`select COUNT(*) as Total from imagelist where %s `, where)
	err := db.GetDB().Sql(sql).Find(&ad)
	_, err = db.GetDB().Sql(totalsql).Get(&total)
	if err != nil {
		panic(err)
	}
	return ad, total.Total
}

func GetFindSort(Module int) int64 {

	type T struct {
		Total int64
	}
	var total T
	where := "1=1"
	if Module != 0 {
		where += fmt.Sprintf(" AND Module = %d ", Module)
	}
	sql := fmt.Sprintf("select Max(Sort) as Total from imagelist where %s group by Module", where)
	_, err := db.GetDB().Sql(sql).Get(&total)
	if err != nil {
		panic(err)
	}
	return total.Total
}

func (c *Imagelist) Insert() error {
	log.Println(c)
	sql := fmt.Sprintf(`INSERT INTO imagelist (Title, Content,Url,IsShow,Module,Sort) VALUES ('%s','%s','%s',%d,%d,%d)`, c.Title, c.Content, c.Url, c.IsShow, c.Module, c.Sort)
	_, err := db.GetDB().Exec(sql)
	if err != nil {
		panic(err)
	}
	return err
}

func (c *Imagelist) Update(col string) error {
	_, err := db.GetDB().Where("ID = ?", c.ID).Cols(col).Update(c)

	return err
}

func (c *Imagelist) Delete() error {
	_, err := db.GetDB().Where("ID = ?", c.ID).Delete(c)
	if err != nil {
		panic(err)
	}
	return err
}
