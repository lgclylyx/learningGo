package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

type MsgTable struct {
	Id          int     `xorm:"'id' INT(11) notnull pk autoincr"`
        Name        string    `xorm:"'name' VARCHAR(30) default ''"`
}


func main() {
	engine, err := xorm.NewEngine("mysql", "root:liuyangxi@(127.0.0.1:3308)/test?charset=utf8")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("success")	
	}
	m := &MsgTable{}
	ok, err := engine.Table("one").Where("id = ?", 1).Get(m)
	if err != nil {
		fmt.Println(err)
	} else {
		if ok {
			fmt.Println("success")
			fmt.Println(*m)
		} else {
			fmt.Println("fail")
		}
	}
	if m.Name == "liuyangxi" {
		m = &MsgTable{
			Name : "wangfengfeng",
		}
	} else {
		m = &MsgTable{
			Name : "liuyangxi",
	    	}
	}
	/*has, err := engine.Table("one").Exist(m)
	if err != nil {
		fmt.Println(err)
	} else {
		if has {
			fmt.Println("exist ", m)
		} else {
			fmt.Println("don't exist", m)
		}
	}*/
	affected, err := engine.Table("one").Id(1).Update(m)
	if err != nil { 
		fmt.Println(err)
	} else {
		fmt.Println("affected num : ", affected)
	}
	
	m = &MsgTable{}	
	ok, err = engine.Sql("select * from one limit 1").Get(m)
	if err == nil && ok {
		fmt.Println("Sql.Get: ", *m)
	} else {
		if ok {
			fmt.Println(err)
		}
	}
	rows, err := engine.Sql("select * from one").Rows(m)
	if err != nil {
		fmt.Println(err)
	} else {
		for rows.Next() {
			if err = rows.Scan(m); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("rows: ",*m)
			}
		}
		rows.Close()
	}
}
