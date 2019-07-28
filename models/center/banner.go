package model

import (
	"fmt"
	db "gmn-admin/db/center"
	"time"
)

type Banner struct {
	Id           int       `json:"Id" xorm:"not null pk autoincr INT(11)"`
	Title        string    `json:"Title" xorm:"not null default '' comment('标题') VARCHAR(255)"`
	Lang         string    `json:"Lang" xorm:"not null comment('语言') VARCHAR(32)"`
	Show         int       `json:"Show" xorm:"not null comment('显示/隐藏') TINYINT(4)"`
	Targetapp    string    `json:"TargetApp" xorm:"not null default '' comment('目标url') VARCHAR(1024)"`
	Targetmobile string    `json:"TargetMobile" xorm:"not null default '' VARCHAR(1024)"`
	Targetpc     string    `json:"TargetPc" xorm:"not null default '' VARCHAR(1024)"`
	Sortapp      int       `json:"SortApp" xorm:"not null comment('排序') INT(11)"`
	Sortmobile   int       `json:"SortMobile" xorm:"not null INT(11)"`
	Sortpc       int       `json:"SortPc" xorm:"not null INT(11)"`
	Imagepc      string    `json:"ImagePc" xorm:"not null default '' comment('图片地址') VARCHAR(1024)"`
	Imageweb     string    `json:"ImageWeb" xorm:"not null default '' VARCHAR(1024)"`
	Imageios     string    `json:"ImageIos" xorm:"not null default '' VARCHAR(1024)"`
	Imageandroid string    `json:"ImageAndroid" xorm:"not null default '' VARCHAR(1024)"`
	Createtime   time.Time `json:"Createtime" xorm:"not null comment('创建时间') DATETIME"`
	Updatetime   time.Time `json:"Updatetime" xorm:"not null comment('更新时间') DATETIME"`
}

func (c *Banner) TableName() string {
	return "Banner"
}

func NewBanner() *Banner {
	return &Banner{
		Createtime: time.Now(),
		Updatetime: time.Now(),
	}
}

func GetBannerById(id int) Banner {
	var ad Banner
	_, err := db.GetDB().Where("Id = ?", id).Get(&ad)
	if err != nil {
		panic(err)
	}
	return ad
}

func GetBanners(offset, pageSize, show int) ([]Banner, int64) {
	var ads = []Banner{}
	// cmd := fmt.Sprintf("SELECT * FROM Banner WHERE `show` = %d	ORDER BY Createtime desc LIMIT %d , %d ", show, ((offset+1)*pageSize - pageSize), pageSize)

	// err := db.GetDB().Sql(cmd).Find(&ads)
	// n, _ := db.GetDB().Sql(cmd).Count(&Banner{})

	var cmd string
	if show != -1 {
		cmd = fmt.Sprintf("`show` = %d", show)
	}
	err := db.GetDB().Where(cmd).Limit(pageSize, offset).OrderBy("Createtime desc").Find(&ads)
	total, _ := db.GetDB().Count(&Banner{})

	if err != nil {
		panic(err)
	}

	return ads, total
}

func (c *Banner) Insert() error {
	_, err := db.GetDB().Insert(c)
	if err != nil {
		panic(err)
	}
	return err
}

func (c *Banner) Update(col string) error {
	_, err := db.GetDB().Where("Id = ?", c.Id).Cols(col).Update(c)
	if err != nil {
		panic(err)
	}
	return err
}

func (c *Banner) Delete() error {
	_, err := db.GetDB().Where("Id = ?", c.Id).Delete(c)
	if err != nil {
		panic(err)
	}
	return err
}
