package main

import (
	"fmt"

	. "github.com/lxn/walk/declarative"
)

func CreateTeacherMenu() {
	mw1 := new(MyMainWindow)
	if err := (MainWindow{
		AssignTo: &mw1.MainWindow,
		Title:    "菜单选择",
		MinSize:  Size{600, 400},
		Size:     Size{600, 400},
		Layout:   VBox{},
		MenuItems: []MenuItem{
			Menu{
				Text: "&教师菜单",
				Items: []MenuItem{
					Separator{},
					Action{
						Text: "&添加学生信息",
						OnTriggered: func() {
							mw1.Close()
							AddStudent()
						},
					},
					Action{
						Text: "&添加学生成绩",
						OnTriggered: func() {
							mw1.Close()
							AddStudentSorce()
						},
					},
					Action{
						Text: "&查询/删除学生成绩",
						OnTriggered: func() {
							mw1.Close()
							Tquery()
						},
					},
					Action{
						Text: "&修改学生成绩",
					},
				},
			},
			Menu{
				Text: "&退出系统",
				Items: []MenuItem{
					Action{
						Text: "退出",
						OnTriggered: func() {
							mw1.Close()
							fmt.Println("Exit system")
						},
					},
				},
			},
		},
	}.Create()); err != nil {

	}
	mw1.Run()
}
