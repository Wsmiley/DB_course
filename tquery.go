package main

import (
	initiator "DB_course/init"
	"DB_course/model"
	"fmt"
	"sort"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

type Condom struct {
	course  string  //课程名
	number  string  //学号
	name    string  //姓名
	sorce   float64 //分数
	checked bool
}
type CondomModel struct {
	walk.TableModelBase
	walk.SorterBase
	sortColumn int
	sortOrder  walk.SortOrder
	items      []*Condom
}

type MyMainWindow3 struct {
	MyMainWindow
	comCA, comCB, comCC *walk.ComboBox
	model               *CondomModel
	tv                  *walk.TableView
}

func (m *CondomModel) RowCount() int {
	return len(m.items)
}
func (m *CondomModel) Checked(row int) bool {
	return m.items[row].checked
}

// Called by the TableView when the user toggled the check box of a given row.
func (m *CondomModel) SetChecked(row int, checked bool) error {
	m.items[row].checked = checked

	return nil
}

// Called by the TableView to sort the model.
func (m *CondomModel) Sort(col int, order walk.SortOrder) error {
	m.sortColumn, m.sortOrder = col, order

	sort.SliceStable(m.items, func(i, j int) bool {

		// c := func(ls bool) bool {
		// 	if m.sortOrder == walk.SortAscending {
		// 		return ls
		// 	}

		// 	return !ls
		// }

		switch m.sortColumn {
		case 0:
			// return c(a.Index < b.Index)

		case 1:
			// return c(a.Bar < b.Bar)

		case 2:
			// return c(a.Baz < b.Baz)

		case 3:
			// return c(a.Quux.Before(b.Quux))
		}

		panic("unreachable")
	})

	return m.SorterBase.Sort(col, order)
}
func NewCondomModel() *CondomModel {
	m := new(CondomModel)

	m.RowsReset()

	return m
}

// Called by the TableView when it needs the text to display for a given cell.
func (m *CondomModel) Value(row, col int) interface{} {
	item := m.items[row]

	switch col {
	case 0:
		return item.course

	case 1:
		return item.number

	case 2:
		return item.name

	case 3:
		return item.sorce
	}

	panic("unexpected col")
}
func Tquery() {
	// mw1 := new(MyMainWindow3)
	mw1 := &MyMainWindow3{model: NewCondomModel()}
	department := make([]string, 2)
	department = []string{"计算机院", "自动化院"}

	var tmp walk.Form

	var courses []model.Course
	var students []model.Student

	var score1 []model.Score
	var comdom []*Condom
	if err := (MainWindow{
		AssignTo: &mw1.MainWindow,
		Title:    "学生成绩管理系统",
		Size:     Size{800, 600},
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
						Text: "&查询/修改学生成绩",
						OnTriggered: func() {
							mw1.Close()
							Tquery()
						},
					},
					Action{
						Text: "&删除学生成绩",
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
			Composite{
				Layout: HBox{MarginsZero: true},
				Children: []Widget{
					HSpacer{},
					Label{
						Text: "院系",
					},
					ComboBox{
						AssignTo: &mw1.comCA,
						Editable: false,
						Model:    department,
						OnCurrentIndexChanged: func() {
							if mw1.comCA.Text() == "计算机院" {
								classdata := make([]string, 2)
								classdata = []string{"计算机171", "计算机172"}
								mw1.comCB.SetModel(classdata)
								var courses1 []model.Course
								if dbError := initiator.MSSQL.Table("courses").Find(&courses1).Error; dbError != nil {
									fmt.Println(dbError)
								}
								coursesdata := make([]string, len(courses1))
								for _, data := range courses1 {
									coursesdata = append(coursesdata, data.Name)
								}
								mw1.comCC.SetModel(coursesdata)
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
						Text: "科目",
					},
					ComboBox{
						AssignTo: &mw1.comCC,
						Editable: false,
					},
					PushButton{
						Text: "查询",
						OnClicked: func() {
							if dbError := initiator.MSSQL.Table("students").Where("Classe=?", mw1.comCB.Text()).Find(&students).Error; dbError != nil {
								fmt.Println(dbError)
							}
							if dbError := initiator.MSSQL.Table("courses").Where("Name=?", mw1.comCC.Text()).First(&courses).Error; dbError != nil {
								fmt.Println(dbError)
							}
							comdom = make([]*Condom, len(students))

							var count = 0
							for _, data := range students {
								if dbError := initiator.MSSQL.Table("scores").Where("Cnumber=? AND Snumber=?", courses[0].Cnumber, data.Snumber).Find(&score1).Error; dbError != nil {
									fmt.Println(dbError)

								}
								if score1 != nil {
									comdom[count] = &Condom{
										course: courses[0].Name,
										name:   data.Name,
										number: data.Snumber,
										sorce:  score1[0].Score}
									count++
								}

							}
							mw1.model.items = comdom
							mw1.model.PublishRowsReset()
							mw1.tv.SetModel(*mw1.model.items[0])
						},
					},
					PushButton{
						Text: "删除",
						OnClicked: func() {
							walk.MsgBox(tmp, "提示", "删除成功", walk.MsgBoxIconInformation)
						},
					},
				},
			},
			Composite{

				Layout: VBox{},
				Children: []Widget{
					TableView{
						AssignTo:         &mw1.tv,
						CheckBoxes:       true,
						ColumnsOrderable: true,
						MultiSelection:   true,
						Model:            mw1.model,
						Columns: []TableViewColumn{
							{Title: "科目"},
							{Title: "学号"},
							{Title: "姓名"},
							{Title: "成绩"},
						},
					},
				},
			},
		},
	}.Create()); err != nil {
		fmt.Println(err)
	}
	mw1.Run()

}
