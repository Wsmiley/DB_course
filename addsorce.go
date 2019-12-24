package main

import (
	initiator "DB_course/init"
	"DB_course/model"
	"fmt"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

type MyMainWindow2 struct {
	MyMainWindow
	name, num                  *walk.LineEdit
	comCA, comCB, comCC, comCD *walk.ComboBox
	cnumA                      *walk.NumberEdit
}

func (mw1 *MyMainWindow2) changeText() {
	if mw1.comCA.Text() == "计算机院" {
		classdata := make([]string, 2)
		classdata = []string{"计算机171", "计算机172"}
		mw1.comCB.SetModel(classdata)
		return
	}
	if mw1.comCA.Text() == "自动化院" {
		classdata := make([]string, 2)
		classdata = []string{"自动化171", "自动化172"}
		mw1.comCB.SetModel(classdata)
		return
	}
}

func AddStudentSorce() {
	mw1 := new(MyMainWindow2)
	var tmp walk.Form
	department := make([]string, 2)
	department = []string{"计算机院", "自动化院"}

	var students []model.Student
	var courses []model.Course

	if err := (MainWindow{
		AssignTo: &mw1.MainWindow,
		Title:    "添加学生考试成绩",
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
				Text: "院系",
			},
			ComboBox{
				AssignTo: &mw1.comCA,
				Editable: false,
				Model:    department,
				OnCurrentIndexChanged: func() {
					if mw1.comCA.Text() != "" {
						var classdata []model.Class
						if dbError := initiator.MSSQL.Where("Deptment=?", mw1.comCA.Text()).Find(&classdata).Error; dbError != nil {
							fmt.Println(dbError)
						}
						class := make([]string, len(classdata))
						for i, data := range classdata {
							if i == 0 {
								class = append(class[:i], data.ClassName)
							} else {
								class = append(class, data.ClassName)
							}

						}
						mw1.comCB.SetModel(class)
					}
				},
			},
			Label{
				Text: "班级",
			},
			ComboBox{
				AssignTo: &mw1.comCB,
				Editable: false,
				OnCurrentIndexChanged: func() {
					if mw1.comCB.Text() != "" {
						if dbError := initiator.MSSQL.Table("students").Where("Class=?", mw1.comCB.Text()).Find(&students).Error; dbError != nil {
							fmt.Println(dbError)
						}
						classdata := make([]string, len(students))
						for i, data := range students {
							if i == 0 {
								classdata = append(classdata[:i], data.Name)
							} else {
								classdata = append(classdata, data.Name)
							}

						}
						mw1.comCC.SetModel(classdata)

						if dbError := initiator.MSSQL.Table("courses").Where("Deptment=?", mw1.comCA.Text()).Find(&courses).Error; dbError != nil {
							fmt.Println(dbError)
						}
						coursesdata := make([]string, len(courses))
						for i, data := range courses {
							if i == 0 {
								coursesdata = append(coursesdata[:i], data.Name)
							} else {
								coursesdata = append(coursesdata, data.Name)
							}

						}

						mw1.comCD.SetModel(coursesdata)
					}
				},
			},
			Label{
				Text: "姓名",
			},
			ComboBox{
				AssignTo: &mw1.comCC,
				Editable: false,
			},
			Label{
				Text: "课程",
			},
			ComboBox{
				AssignTo: &mw1.comCD,
				Editable: false,
			},
			Label{
				Text: "成绩",
			},
			NumberEdit{
				AssignTo: &mw1.cnumA,
				Decimals: 2,
				MaxValue: 150,
				MinValue: 0,
			},
			PushButton{
				Text:    "提交",
				MinSize: Size{120, 50},
				OnClicked: func() {
					if mw1.cnumA.Value() == 0 {
						walk.MsgBox(tmp, "提示", "成绩不能为0", walk.MsgBoxIconInformation)
						return
					}
					var score = model.Score{
						Snumber: students[mw1.comCC.CurrentIndex()].Snumber,
						Cnumber: courses[mw1.comCD.CurrentIndex()].Cnumber,
						Score:   mw1.cnumA.Value()}

					initiator.MSSQL.Create(&score)
					walk.MsgBox(tmp, "提示", "提交成功", walk.MsgBoxIconInformation)
					fmt.Println(" successfully added")
					mw1.Close()
					AddStudentSorce()
					return

				},
			},
		},
	}.Create()); err != nil {

	}
	mw1.Run()
}
