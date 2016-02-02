package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"     //引入beego的orm
	_ "github.com/go-sql-driver/mysql" //引入beego的mysql驱动
)

type GoBar struct {
	Id      int
	BarId   int64
	BarName string `orm:"size(60)"`
	BarDesc string `orm:"size(100)"`
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("go", "mysql", "cmsuser:cms@tcp(10.255.223.241:3306)/bookbar_fuzuokui?charset=utf8")
	orm.SetMaxIdleConns("go", 30) //设置数据库最大空闲连接
	orm.SetMaxOpenConns("go", 30) //设置数据库最大连接数
	orm.RegisterModel(new(GoBar))
}

func (goBar *GoBar) Add(bars []Bar) (int, error) {
	o := orm.NewOrm()
	o.Using("go")
	var newBars []GoBar = make([]GoBar, len(bars))
	for i, bar := range bars {
		newBars[i] = GoBar{BarId: bar.BarId, BarName: bar.BarName, BarDesc: bar.BarDesc}
	}
	_, err := o.InsertMulti(len(newBars), newBars)
	if err != nil {
		fmt.Println(err)
	}
	return len(newBars), err
}
