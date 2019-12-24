package initiator

import (
	"DB_course/model"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

// db.SingularTable(true) // 如果设置为true,`User`的默认表名为`user`,使用`TableName`设置的表名不受影响
var MSSQL *gorm.DB

func createTableAdmins(db *gorm.DB) {
	db.CreateTable(&model.Admins{})
}

func createTableStudent(db *gorm.DB) {
	db.CreateTable(&model.Student{})
}

func createTableScore(db *gorm.DB) {
	db.CreateTable(&model.Score{})
}

func createTableCourse(db *gorm.DB) {
	db.CreateTable(&model.Course{})
}
func createTableClass(db *gorm.DB) {
	db.CreateTable(&model.Class{})
}
func createTableStudentCounts(db *gorm.DB) {
	db.CreateTable(&model.StudentCount{})
}

func init() {
	db, err := gorm.Open("mssql", "sqlserver://SA:QQcc7711.@localhost:1433?database=DB_course")
	if err != nil {
		panic("Database connection failed")
	}
	fmt.Println("Database connection success")
	if !db.HasTable("admins") {
		createTableAdmins(db)
		fmt.Println("create Admin Table")
	}
	if !db.HasTable("students") {
		createTableStudent(db)
		fmt.Println("create Student Table")
	}
	if !db.HasTable("courses") {
		createTableCourse(db)
		fmt.Println("create Course Table")
	}
	if !db.HasTable("scores") {
		createTableScore(db)
		fmt.Println("create Score Table")
	}
	if !db.HasTable("classes") {
		createTableClass(db)
		fmt.Println("create class Table")
	}
	if !db.HasTable("student_counts") {
		createTableStudentCounts(db)
		fmt.Println("create studentcounts Table")
	}

	// courses1 := model.Course{Cnumber: "000001", Name: "数据库", Deptment: "计算机院"}
	// db.Create(&courses1)
	// courses2 := model.Course{Cnumber: "000002", Name: "操作系统", Deptment: "计算机院"}
	// db.Create(&courses2)
	// courses3 := model.Course{Cnumber: "000003", Name: "微机原理", Deptment: "计算机院"}
	// db.Create(&courses3)
	// courses4 := model.Course{Cnumber: "000004", Name: "ARM嵌入式", Deptment: "计算机院"}
	// db.Create(&courses4)
	// class1 := model.Class{ClassId: "20190001", ClassName: "计算机171", Deptment: "计算机院"}
	// db.Create(&class1)
	// class2 := model.Class{ClassId: "20190002", ClassName: "计算机172", Deptment: "计算机院"}
	// db.Create(&class2)
	// class3 := model.Class{ClassId: "20190003", ClassName: "自动化171", Deptment: "自动化院"}
	// db.Create(&class3)
	// class4 := model.Class{ClassId: "20190004", ClassName: "自动化172", Deptment: "自动化院"}
	// db.Create(&class4)
	// courses1 := model.Course{Cnumber: "000005", Name: "自动控制原理", Deptment: "自动化院"}
	// db.Create(&courses1)
	// courses2 := model.Course{Cnumber: "000006", Name: "电路", Deptment: "自动化院"}
	// db.Create(&courses2)
	// courses3 := model.Course{Cnumber: "000007", Name: "模拟电子技术", Deptment: "自动化院"}
	// db.Create(&courses3)
	// courses4 := model.Course{Cnumber: "000008", Name: "单片机与嵌入式系统原理", Deptment: "自动化院"}
	// db.Create(&courses4)
	MSSQL = db

}

//docker run -e "ACCEPT_EULA=Y" -e "SA_PASSWORD=QQcc7711." -p 1433:1433 --name sql1 -d mcr.microsoft.com/mssql/server:2017-latest
// -- docker exec -it sql1 "bash"
// --
