package db

import (
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var mysql *xorm.Engine

func InitDatabase(host string, user string, pass string, dbname string) {
	connString := user + ":" + pass + "@tcp(" + host + ")/" + dbname + "?charset=utf8"
	Init(connString)
}

func Init(conn string) {
	var err error
	mysql, err = xorm.NewEngine("mysql", conn)
	if err != nil {
		log.Fatal("mysql connect err: ", err, "\n")
	}
	mysql.SetMaxConns(200)
	mysql.SetMaxOpenConns(180)
	//mysql.SetLogger(xorm.NewSimpleLogger(os.Stderr))
	//mysql.ShowSQL(true)
}

func GetDB() *xorm.Engine {
	return mysql
}

func Debug() {
	mysql.SetLogger(xorm.NewSimpleLogger(os.Stderr))
	mysql.ShowSQL(true)
}

type Trasaction struct {
	session *xorm.Session
}

func NewTrasaction() *Trasaction {
	return &Trasaction{
		session: GetDB().NewSession(),
	}
}

func (c *Trasaction) Inst() *xorm.Session {
	return c.session
}

func (c *Trasaction) Begin() {
	c.session.Begin()
}

func (c *Trasaction) Commit() {
	c.session.Commit()
}
