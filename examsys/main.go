package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Student struct {
	ID        int       `gorm:"primaryKey:true;autoIncrement:true;column:id"`
	Name      string    `gorm:"column:name"`
	Age       int       `gorm:"column:age"`
	Addr      string    `gorm:"column:addr"`
	Examtype  string    `gorm:"column:examtype"`
	Marksheet Marksheet `gorm:"foreignkey:student_id;references:id"`
	// Marksheet Marksheet `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
type Marksheet struct {
	Id int `gorm:"primaryKey:true;autoIncrement:true;column:id"`
	// Student_id int    `gorm:"foreignKey:id"`
	Student_id int    `gorm:"column:student_id"`
	Subject    string `gorm:"column:subject"`
	Marks      int    `gorm:"column:marks"`
	//Student    Student `"foreignkey:id;references:student_id"`
}

func main() {
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "1213"
		dbname   = "db_1"
	)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d",
		host, user, password, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	} else {
		fmt.Println("Db Connected...")
	}

	db.AutoMigrate(&Student{})

	db.AutoMigrate(&Marksheet{})

	// db.Model(&student{}).

	//db.Model(&marksheet{}).("student_id", "student(id)", "RESTRICT", "RESTRICT")
	// db.Model(&marksheet{}).AddForeignKey("u_id", "t_user(id)", "RESTRICT", "RESTRICT")
	chk := true

	for chk == true {

		fmt.Println("Enter 1 to add data to student table..")
		fmt.Println("Enter 2 to add data to marksheet table..")
		fmt.Println("Enter 3 for Searching..")
		fmt.Println("Enter 4 to exit...")

		var ch int

		fmt.Scan(&ch)

		switch ch {
		case 1:
			fmt.Println("Enter name :")
			var nm string
			fmt.Scan(&nm)

			fmt.Println("Enter age :")
			var ag int
			fmt.Scan(&ag)

			fmt.Println("Enter Address :")
			var addr string
			fmt.Scan(&addr)

			fmt.Println("Enter EmamType :")
			var et string
			fmt.Scan(&et)

			fmt.Println("Enter subject :")
			var sub string
			fmt.Scan(&sub)

			fmt.Println("Enter Subject Marks:")
			var tm int
			fmt.Scan(&tm)

			fmt.Println("Enter Student_id:")
			var stu int
			fmt.Scan(&stu)

			u := Student{Name: nm, Age: ag, Addr: addr, Examtype: et}
			db.Create(&u)
			m := Marksheet{Subject: sub, Marks: tm, Student_id: stu}
			db.Create(&m)

		case 2:
			// fmt.Println("Enter Student id :")
			// var sid int
			// fmt.Scan(&sid)

			fmt.Println("Enter subject :")
			var sub string
			fmt.Scan(&sub)

			fmt.Println("Enter Student Marks:")
			var tm int
			fmt.Scan(&tm)

			mk := Marksheet{Subject: sub, Marks: tm}
			db.Create(&mk)

		case 3:
			// fmt.Println("Enter name :")
			// var s string
			// fmt.Scan(&s)
			// // std := []Student{}

			// // mark := Marksheet{}
			// // mark.Id = mark.Student_id

			// var std []Student

			// // Name: s
			// db.Where(&Student{Name: s}).Find(&std)

			// for _, v := range std {
			// 	fmt.Printf("ID: '%v' Name: '%v' Age: '%v' ExamType: '%v' Subject: %v Marks %v \n",
			// 		v.ID, v.Name, v.Age, v.Addr, v.Examtype, v.Subject, v.Marks)
			// }
		case 4:
			chk = false

		}

	}

}
