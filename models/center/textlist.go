package model

import (
	"fmt"
	db "gmn-admin/db/center"
	"time"
)

type Textlist struct {
	ID         int       `json:"ID" xorm:"ID"`
	Content    string    `json:"Content" xorm:"Content"`
	Module     int       `json:"Module" xorm:"Module"`
	Createtime time.Time `json:"CreateTime" xorm:"CreateTime"`
	Updatetime time.Time `json:"UpdateTime" xorm:"UpdateTime"`
}

func NewTextlist() *Textlist {
	return &Textlist{}
}

func GetTextlist() []Textlist {
	var ad []Textlist

	sql := fmt.Sprintf("select * from textlist")
	err := db.GetDB().Sql(sql).Find(&ad)
	if err != nil {
		panic(err)
	}
	return ad
}

func (c *Textlist) Insert() error {
	sql := fmt.Sprintf(`INSERT INTO textlist(Content,Module) VALUES("%s", %d) ON DUPLICATE KEY UPDATE Content=VALUES(Content),Module=VALUES(Module)`, c.Content, c.Module)
	_, err := db.GetDB().Exec(sql)
	if err != nil {
		panic(err)
	}
	return err
}
