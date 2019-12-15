package main

import (
	"fmt"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func Select() {
	walk.Resources.SetRootDirPath("./img")
	mw := new(MyMainWindow)

	if err := (MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    "学生成绩管理系统",
		MinSize:  Size{270, 290},
		Size:     Size{270, 290},
		Layout:   VBox{},
		Children: []Widget{
			HSpacer{},
			ImageView{
				Image:  "njit.png",
				Margin: 10,
			},
			Composite{
				Layout: VBox{},
				Children: []Widget{
					PushButton{
						Text:    "教师登入",
						MinSize: Size{120, 50},
						OnClicked: func() {
							fmt.Println("Select successful")
							mw.Close()
							Tlogin()

						},
					},
					PushButton{
						Text:    "学生登入",
						MinSize: Size{120, 50},
						OnClicked: func() {
							fmt.Println("Select successful")
							mw.Close()
							Slogin()

						},
					},
				},
			},
		},
	}.Create()); err != nil {

	}
	mw.Run()
}
