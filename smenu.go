package main

import (
	"fmt"

	. "github.com/lxn/walk/declarative"
)

func CreateStudentMenu(studentnum string) {
	mw1 := new(MyMainWindow)
	if err := (MainWindow{
		AssignTo: &mw1.MainWindow,
		Title:    "菜单选择",
		MinSize:  Size{600, 400},
		Size:     Size{600, 400},
		Layout:   VBox{},
		MenuItems: []MenuItem{
			Menu{
				Text: "&学生菜单",
				Items: []MenuItem{
					Separator{},
					Action{
						Text: "&查询成绩",
						OnTriggered: func() {
							mw1.Close()
							Squery("202170109")
						},
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
