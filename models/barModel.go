package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"     //引入beego的orm
	_ "github.com/go-sql-driver/mysql" //引入beego的mysql驱动
	"io"
	"os"
	"time"
)

type Bar struct {
	BarId              int64 `orm:"pk;auto"`
	BarName            string
	BarDesc            string
	BarImgUrl          string
	bar_type           int
	bar_status         int
	member_num         int
	article_num        int
	creator            int64
	create_date        time.Time
	last_modified_user int64
	last_modified_date time.Time
	remark             string
	isActivity         int
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "cmsuser:cms@tcp(10.255.223.241:3306)/bookbar?charset=utf8")
	orm.SetMaxIdleConns("default", 30) //设置数据库最大空闲连接
	orm.SetMaxOpenConns("default", 30) //设置数据库最大连接数
	orm.RegisterModel(new(Bar))
	orm.Debug = true
}

func (bar *Bar) QueryAddBars(barId int64) ([]Bar, error) {
	o := orm.NewOrm()
	var data []Bar
	_, err := o.QueryTable("bar").Filter("BarId__gt", barId).Limit(10).All(&data)
	if err != nil {
		fmt.Println(err)
	}
	var w io.Writer = os.Stdout
	orm.DebugLog = orm.NewLog(w)
	return data, err
}
