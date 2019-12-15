package main

import (
	"fmt"

	. "github.com/lxn/walk/declarative"
)

func Select() {
	// var tmp walk.Form
	mw := new(MyMainWindow)
	if err := (MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    "学生成绩管理系统",
		MinSize:  Size{270, 290},
		Size:     Size{270, 290},
		Layout:   VBox{},
		Children: []Widget{

			PushButton{
				Text:    "教师登入",
				MinSize: Size{120, 50},
				OnClicked: func() {
					mw.Close()
					Tlogin()
					fmt.Println("Select successful")
				},
			},
			PushButton{
				Text:    "学生登入",
				MinSize: Size{120, 50},
				OnClicked: func() {
					mw.Close()
					Slogin()
					fmt.Println("Select successful")

				},
			},
		},
	}.Create()); err != nil {

	}
	mw.Run()
}
