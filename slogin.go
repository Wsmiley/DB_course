package main

import (
	"fmt"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func Slogin() {
	mw := new(MyMainWindow)
	var tmp walk.Form
	var usernameTE, passwordTE *walk.LineEdit
	if err := (MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    "Student grade management system",
		MinSize:  Size{270, 290},
		Layout:   VBox{},
		Children: []Widget{
			Composite{
				Layout: VBox{},
				Children: []Widget{
					VSplitter{
						Children: []Widget{
							Label{
								Text: "用户名",
							},
							LineEdit{
								MinSize:  Size{160, 0},
								AssignTo: &usernameTE,
							},
						},
					},
					VSplitter{
						Children: []Widget{
							Label{MaxSize: Size{160, 40},
								Text: "密码",
							},
							LineEdit{
								MinSize:  Size{160, 0},
								AssignTo: &passwordTE,
							},
						},
					},
				},
			},

			PushButton{
				Text:    "登录",
				MinSize: Size{120, 50},
				OnClicked: func() {
					if usernameTE.Text() == "" {

						walk.MsgBox(tmp, "警告", "用户名为空", walk.MsgBoxIconInformation)
						return
					}
					if passwordTE.Text() == "" {
						walk.MsgBox(tmp, "警告", "密码为空", walk.MsgBoxIconInformation)
						return
					}

					var tmp walk.Form
					walk.MsgBox(tmp, "提示", "欢迎", walk.MsgBoxIconInformation)
					fmt.Println("Login successful")
					mw.Close()
					CreateStudentMenu()
				},
			},
		},
	}.Create()); err != nil {

	}
	mw.Run()
}