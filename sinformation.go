package main

import (
	initiator "DB_course/init"
	"DB_course/model"
	"fmt"

	. "github.com/lxn/walk/declarative"
)

func Sinformation(studentnum string) {
	mw1 := new(MyMainWindow)

	var students []model.Student
	if dbError := initiator.MSSQL.Table("students").Where("Snumber=?", studentnum).Find(&students).Error; dbError != nil {
		fmt.Println(dbError)
	}
	if err := (MainWindow{
		AssignTo: &mw1.MainWindow,
		Title:    "个人信息",
		MinSize:  Size{200, 200},
		Size:     Size{250, 400},
		Layout:   VBox{},
		Children: []Widget{
			Composite{
				Layout: VBox{},
				Children: []Widget{
					VSplitter{
						Children: []Widget{
							Label{
								Text: "姓名",
							},
							LineEdit{
								Text:     students[0].Name,
								ReadOnly: true,
							},
							Label{
								Text: "学号",
							},
							LineEdit{
								Text:     studentnum,
								ReadOnly: true,
							},
							Label{
								Text: "性别",
							},
							LineEdit{
								Text:     students[0].Sex,
								ReadOnly: true,
							},
							Label{
								Text: "班级",
							},
							LineEdit{
								Text:     students[0].Class,
								ReadOnly: true,
							},
							Label{
								Text: "院系",
							},
							LineEdit{
								Text:     students[0].Dept,
								ReadOnly: true,
							},
						},
					},
				},
			},
		},
	}.Create()); err != nil {

	}
	mw1.Run()
}
