package main

import (
	initiator "DB_course/init"
	"DB_course/model"
	"fmt"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

type CondomStudent struct {
	course string  //课程名
	sorce  float64 //分数
	// cnumber string  //课程号
	checked bool
}
type CondomStudentModel struct {
	walk.TableModelBase
	walk.SorterBase
	sortColumn int
	sortOrder  walk.SortOrder
	items      []*CondomStudent
}

type MyMainWindow4 struct {
	MyMainWindow
	comCA, comCB, comCC *walk.ComboBox
	model               *CondomStudentModel
	tv                  *walk.TableView
	lA, lB              *walk.Label
}

func (m *CondomStudentModel) RowCount() int {
	return len(m.items)
}

// Called by the TableView to sort the model.
func (m *CondomStudentModel) Sort(col int, order walk.SortOrder) error {
	m.sortColumn, m.sortOrder = col, order
	return m.SorterBase.Sort(col, order)
}

func (m *CondomStudentModel) Less(i, j int) bool {
	a, b := m.items[i], m.items[j]

	c := func(ls bool) bool {
		if m.sortOrder == walk.SortAscending {
			return ls
		}

		return !ls
	}

	switch m.sortColumn {
	case 0:
		return c(a.course < b.course)
	case 1:
		return c(a.sorce < b.sorce)
	}

	panic("unreachable")
}

func NewCondomStudnetModel() *CondomStudentModel {
	m := new(CondomStudentModel)

	m.RowsReset()

	return m
}

// Called by the TableView when it needs the text to display for a given cell.
func (m *CondomStudentModel) Value(row, col int) interface{} {
	item := m.items[row]
	if item != nil {
		switch col {
		case 0:
			return item.course
		case 1:
			return item.sorce
		}
	} else {
		return ""
	}

	panic("unexpected col")
}
func (m *CondomStudentModel) Swap(i, j int) {
	m.items[i], m.items[j] = m.items[j], m.items[i]
}

func Squery(studentnum string) {
	mw1 := &MyMainWindow4{model: NewCondomStudnetModel()}

	var comdom []*CondomStudent
	var score1 []model.Score
	var courses []model.Course
	if err := (MainWindow{
		AssignTo: &mw1.MainWindow,
		Title:    "学生成绩管理系统",
		Size:     Size{400, 400},
		Layout:   VBox{},
		MenuItems: []MenuItem{
			Menu{
				Text: "&学生菜单",
				Items: []MenuItem{
					Separator{},
					Action{
						Text: "&个人信息",
						OnTriggered: func() {
							Sinformation(studentnum)
						},
					},
					Action{
						Text: "&查询成绩",
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
						AssignTo: &mw1.lA,
						Text:     "",
					},

					Label{
						AssignTo: &mw1.lB,
						Text:     "",
					},
					PushButton{
						Text: "查询成绩",
						OnClicked: func() {
							comdom = make([]*CondomStudent, 20)
							if dbError := initiator.MSSQL.Where("Snumber=?", studentnum).Find(&score1).Error; dbError != nil {
								fmt.Println(dbError)
							}
							if dbError := initiator.MSSQL.Table("courses").Find(&courses).Error; dbError != nil {
								fmt.Println(dbError)
							}
							if score1 != nil {

								for i, data1 := range score1 {
									for _, data2 := range courses {
										if data1.Cnumber == data2.Cnumber {
											comdom[i] = &CondomStudent{
												course: data2.Name,
												sorce:  data1.Score}
											break
										}
									}
								}
							}
							mw1.model.items = comdom
							mw1.model.PublishRowsReset()
							mw1.tv.SetModel(*mw1.model.items[0])
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
							{Title: "成绩"},
						},
						OnItemActivated: mw1.tv_ItemActivated,
					},
				},
			},
		},
	}.Create()); err != nil {
		fmt.Println(err)
	}
	mw1.Run()

}
func (mw *MyMainWindow4) tv_ItemActivated() {
	msg := ``
	for _, i := range mw.tv.SelectedIndexes() {
		msg = msg + "\n" + mw.model.items[i].course
	}
	walk.MsgBox(mw, "title", msg, walk.MsgBoxIconInformation)
}
