package main

import (
	initiator "DB_course/init"
	"DB_course/model"
	"fmt"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

type MyMainWindow1 struct {
	MyMainWindow
	name, num           *walk.LineEdit
	comCA, comCB, comCC *walk.ComboBox
}

func AddStudent() {
	mw1 := new(MyMainWindow1)
	var tmp walk.Form

	department := make([]string, 1)
	department = []string{"计算机院", "自动化院"}

	sexdata := make([]string, 2)
	sexdata = []string{"男", "女"}

	if err := (MainWindow{
		AssignTo: &mw1.MainWindow,
		Title:    "添加学生信息",
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
		Children: []Widget{
			Label{
				Text: "姓名",
			},
			LineEdit{
				MinSize:  Size{160, 0},
				AssignTo: &mw1.name,
			},
			Label{
				Text: "学号",
			},
			LineEdit{
				MinSize:  Size{160, 0},
				AssignTo: &mw1.num,
			},
			Label{
				Text: "院系",
			},
			ComboBox{
				AssignTo: &mw1.comCA,
				Editable: false,
				Model:    department,
				OnCurrentIndexChanged: func() {
					//因从数据库中提取数据，再做修改
					if mw1.comCA.Text() == "计算机院" {
						classdata := make([]string, 2)
						classdata = []string{"计算机171", "计算机172"}
						mw1.comCB.SetModel(classdata)
					}
					if mw1.comCA.Text() == "自动化院" {
						classdata := make([]string, 2)
						classdata = []string{"自动化171", "自动化172"}
						mw1.comCB.SetModel(classdata)
					}
				},
			},
			Label{
				Text: "班级",
			},
			ComboBox{
				AssignTo: &mw1.comCB,
				Editable: false,
			},
			Label{
				Text: "性别",
			},
			ComboBox{
				AssignTo: &mw1.comCC,
				Editable: false,
				Model:    sexdata,
			},

			PushButton{
				Text:    "提交",
				MinSize: Size{120, 50},
				OnClicked: func() {
					if mw1.name.Text() == "" {
						walk.MsgBox(tmp, "警告", "姓名为空", walk.MsgBoxIconInformation)
						return
					}
					if mw1.num.Text() == "" {
						walk.MsgBox(tmp, "警告", "学号为空", walk.MsgBoxIconInformation)
						return
					}
					if mw1.comCA.Text() == "" {
						walk.MsgBox(tmp, "警告", "班级未选", walk.MsgBoxIconInformation)
						return
					}
					if mw1.comCB.Text() == "" {
						walk.MsgBox(tmp, "警告", "院系未选", walk.MsgBoxIconInformation)
						return
					}
					if mw1.comCC.Text() == "" {
						walk.MsgBox(tmp, "警告", "性别未选", walk.MsgBoxIconInformation)
						return
					}
					var student = model.Student{
						Snumber: mw1.num.Text(),
						Name:    mw1.name.Text(),
						Classe:  mw1.comCB.Text(),
						Sex:     mw1.comCC.Text(),
						Dept:    mw1.comCA.Text()}
					initiator.MSSQL.Create(&student)
					mw1.Close()
					walk.MsgBox(tmp, "提示", "提交成功", walk.MsgBoxIconInformation)
					fmt.Println(" successfully added")
					AddStudent()

				},
			},
		},
	}.Create()); err != nil {

	}
	mw1.Run()
}
