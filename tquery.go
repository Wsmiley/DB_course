package main

import (
	initiator "DB_course/init"
	"DB_course/model"
	"fmt"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

type Condom struct {
	course  string  //课程名
	number  string  //学号
	name    string  //姓名
	sorce   float64 //分数
	cnumber string  //课程号

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

// Called by the TableView to sort the model.
func (m *CondomModel) Sort(col int, order walk.SortOrder) error {
	m.sortColumn, m.sortOrder = col, order
	// sort.Stable(m)
	return m.SorterBase.Sort(col, order)
}

func (m *CondomModel) Less(i, j int) bool {
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
		return c(a.name < b.name)

	case 2:
		return c(a.sorce < b.sorce)
	}

	panic("unreachable")
}

func NewCondomModel() *CondomModel {
	m := new(CondomModel)

	m.RowsReset()

	return m
}

// Called by the TableView when it needs the text to display for a given cell.
func (m *CondomModel) Value(row, col int) interface{} {
	item := m.items[row]
	if item != nil {
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
	} else {
		return ""
	}

	panic("unexpected col")
}

func (m *CondomModel) Swap(i, j int) {
	m.items[i], m.items[j] = m.items[j], m.items[i]
}

func Tquery() {
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
								var courses1 []model.Course
								if dbError := initiator.MSSQL.Table("courses").Where("Deptment=?", mw1.comCA.Text()).Find(&courses1).Error; dbError != nil {
									fmt.Println(dbError)
								}
								coursesdata := make([]string, len(courses1))
								for i, data := range courses1 {
									if i == 0 {
										coursesdata = append(coursesdata[:1], data.Name)
									} else {
										coursesdata = append(coursesdata, data.Name)
									}
								}
								mw1.comCC.SetModel(coursesdata)
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
						Text: "查询成绩",
						OnClicked: func() {
							if mw1.comCA.Text() == "" {
								walk.MsgBox(tmp, "警告", "请选择院系", walk.MsgBoxIconInformation)
								return
							}
							if mw1.comCB.Text() == "" {
								walk.MsgBox(tmp, "警告", "请选择班级", walk.MsgBoxIconInformation)
								return
							}
							if mw1.comCC.Text() != "" {
								if dbError := initiator.MSSQL.Table("students").Where("Class=?", mw1.comCB.Text()).Find(&students).Error; dbError != nil {
									fmt.Println(dbError)
								}
								if dbError := initiator.MSSQL.Table("courses").Where("Name=?", mw1.comCC.Text()).First(&courses).Error; dbError != nil {
									fmt.Println(dbError)
								}
								comdom = make([]*Condom, len(students))
								var count = 0
								for _, data := range students {
									dbError := initiator.MSSQL.Where("Cnumber=? AND Snumber=?", courses[0].Cnumber, data.Snumber).Find(&score1).Error
									if dbError != nil {
										fmt.Println(dbError)
									}
									if score1 != nil {
										if len(score1) > 0 {
											comdom[count] = &Condom{
												course:  courses[0].Name,
												name:    data.Name,
												number:  data.Snumber,
												sorce:   score1[0].Score,
												cnumber: courses[0].Cnumber}
											count++
										}
									}
								}
								mw1.model.items = comdom
								mw1.model.PublishRowsReset()
								mw1.tv.SetModel(*mw1.model.items[0])
							} else {
								if dbError := initiator.MSSQL.Table("students").Where("Class=?", mw1.comCB.Text()).Find(&students).Error; dbError != nil {
									fmt.Println(dbError)
								}
								if dbError := initiator.MSSQL.Table("courses").Find(&courses).Error; dbError != nil {
									fmt.Println(dbError)
								}
								comdom = make([]*Condom, 100)
								var count = 0
								for _, data := range students {
									dbError := initiator.MSSQL.Where("Snumber=?", data.Snumber).Find(&score1).Error
									if dbError != nil {
										fmt.Println(dbError)
									}
									if score1 != nil {
										for _, data1 := range score1 {
											num := 0
											for i, data2 := range courses {
												if data1.Cnumber == data2.Cnumber {
													num = i
													break
												}
											}
											comdom[count] = &Condom{
												course:  courses[num].Name,
												name:    data.Name,
												number:  data.Snumber,
												sorce:   data1.Score,
												cnumber: data1.Cnumber}
											count++
										}
									}
								}
								mw1.model.items = comdom
								mw1.model.PublishRowsReset()
								mw1.tv.SetModel(*mw1.model.items[0])
							}
						},
					},
					PushButton{
						Text: "删除",
						OnClicked: func() {
							items := []*Condom{}
							remove := mw1.tv.SelectedIndexes()
							for i, x := range mw1.model.items {
								remove_ok := false
								for _, j := range remove {
									if i == j {
										remove_ok = true
										if dbError := initiator.MSSQL.Table("scores").Where("Snumber=? AND Cnumber=? ", x.number, x.cnumber).Delete(model.Score{}).Error; dbError != nil {
											fmt.Println(dbError)
										}
									}
								}
								if !remove_ok {
									items = append(items, x)
								}
							}
							mw1.model.items = items
							mw1.model.PublishRowsReset()
							mw1.tv.SetSelectedIndexes([]int{})
							walk.MsgBox(tmp, "警告", "删除成功", walk.MsgBoxIconInformation)

						},
					},
					PushButton{
						Text: "创建Excel",
						OnClicked: func() {
							CreateExcel(mw1.model.items)
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
func (mw *MyMainWindow3) tv_ItemActivated() {
	msg := ``
	for _, i := range mw.tv.SelectedIndexes() {
		msg = msg + "\n" + mw.model.items[i].name
	}
	walk.MsgBox(mw, "title", msg, walk.MsgBoxIconInformation)
}
