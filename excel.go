package main

import (
	"fmt"
	"strconv"

	"github.com/tealeg/xlsx"
)

func CreateExcel(c []*Condom) {

	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row, row1 *xlsx.Row
	var cell *xlsx.Cell
	var err error

	file = xlsx.NewFile()
	sheet, err = file.AddSheet("Sheet1")
	if err != nil {
		fmt.Printf(err.Error())
	}
	row = sheet.AddRow()
	row.SetHeightCM(1)
	cell = row.AddCell()
	cell.Value = "学号"
	cell = row.AddCell()
	cell.Value = "姓名"
	cell = row.AddCell()
	cell.Value = "课程"
	cell = row.AddCell()
	cell.Value = "成绩"

	for _, data := range c {
		if data == nil {
			break
		}
		fmt.Println(data)
		row1 = sheet.AddRow()
		row1.SetHeightCM(1)
		cell = row1.AddCell()
		cell.Value = data.number
		cell = row1.AddCell()
		cell.Value = data.name
		cell = row1.AddCell()
		cell.Value = data.course
		cell = row1.AddCell()
		cell.Value = strconv.FormatFloat(data.sorce, 'f', -1, 64)
	}

	err = file.Save("学生成绩.xlsx")
	if err != nil {
		fmt.Printf(err.Error())
	}
}
