package model

import (
	"gmn-admin/common"
	time "gmn-admin/common"
	db "gmn-admin/db/center"
)

type Admin struct {
	Uid           int       `json:"uid" xorm:"not null pk autoincr INT(11)"`
	Username      string    `json:"username" xorm:"not null VARCHAR(255)"`
	Password      string    `json:"password" xorm:"not null VARCHAR(255)"`
	Crtime        time.Time `json:"crtime" xorm:"not null DATETIME"`
	Uptime        time.Time `json:"uptime" xorm:"not null DATETIME"`
	LastLoginIp   string    `json:"last_login_ip" xorm:"not null VARCHAR(255)"`
	LastLoginTime time.Time `json:"last_login_time" xorm:"not null DATETIME"`
}

func (c *Admin) TableName() string {
	return common.TableName("admin")
}

func NewAdmin(username string, password, salt string) *Admin {
	return &Admin{
		Username: username,
		Password: password,
		Crtime:   time.Now(),
		Uptime:   time.Now(),
	}
}

func GetAdmins() []Admin {
	var admins = []Admin{}
	err := db.GetDB().Where("state != 1").Find(&admins)
	if err != nil {
		panic(err)
	}
	return admins
}

func GetAdminById(uid int) Admin {
	var admin Admin
	_, err := db.GetDB().Where("uid = ? ", uid).Get(&admin)
	if err != nil {
		panic(err)
	}
	return admin
}

func GetAdminByUsername(username string) Admin {
	var admin Admin
	_, err := db.GetDB().Where("username = ?", username).Get(&admin)
	if err != nil {
		panic(err)
	}
	return admin
}

func (c *Admin) Update(col string) {
	_, err := db.GetDB().Cols(col).Where("uid = ?", c.Uid).Update(c)
	if err != nil {
		panic(err)
	}
}

func (c *Admin) Insert() {
	_, err := db.GetDB().Insert(c)
	if err != nil {
		panic(err)
	}
}
