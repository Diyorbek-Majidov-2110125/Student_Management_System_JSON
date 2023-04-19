package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"Student_Management/storage/repo"
	"Student_Management/storage/jsons"
	"bufio"
	"os"
)



func main() {

	var studentList jsons.Students

	data, err := ioutil.ReadFile("../storage/jsons/student.json")
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}

	err = json.Unmarshal(data, &studentList)
	if err != nil {
		fmt.Println("Error in Unmarshaling:", err)
		return
	}

	fmt.Println("-------------------------------------------------------------------")
	fmt.Println("--------------Welcome to Student Management System ----------------")
	fmt.Println("-------------------------------------------------------------------")
	var choice int
	for choice != 6 {
		fmt.Println("-------------------------------------------------------------------")
		fmt.Println("------------------------Enter your choice:-------------------------")
		fmt.Println("-------------------------------------------------------------------")
		fmt.Println("<1> - See the list of students")
		fmt.Println("<2> - Get the student by ID")
		fmt.Println("<3> - Add the student")
		fmt.Println("<4> - Remove the student by ID")
		fmt.Println("<5> - Update the student's information")
		fmt.Println("<6> - Exit")

		fmt.Println("\nEnter your choice:")
		fmt.Scanln(&choice)

		switch choice {
			case 1:	
				students := studentList.GetStudents()
				for _, student := range students {
fmt.Println("--------------------------------------------------------------------------")
fmt.Printf(`
Student Id: %v
Name: %v
Overall Grade: %v
Section: %v
Course: %v
Subject: %v
Marks: %v
Subject: %v
Marks: %v
Subject: %v
Marks: %v
`,student.ID,student.Name,student.Grade,student.Section,student.Course,student.Subjects[0].Name,student.Subjects[0].Marks,student.Subjects[1].Name,student.Subjects[1].Marks,student.Subjects[2].Name,student.Subjects[2].Marks)
}
			case 2: 
				var id int
				fmt.Println("Enter the student ID")
				fmt.Scanln(&id)
				student := studentList.GetStudentByID(id)
				fmt.Println("--------------------------------------------------------------------------")
fmt.Printf(`
Student Id: %v
Name: %v
Overall Grade: %v
Section: %v
Course: %v
Subject: %v
Marks: %v
Subject: %v
Marks: %v
Subject: %v
Marks: %v
`,student.ID,student.Name,student.Grade,student.Section,student.Course,student.Subjects[0].Name,student.Subjects[0].Marks,student.Course,student.Subjects[1].Name,student.Subjects[1].Marks,student.Course,student.Subjects[2].Name,student.Subjects[2].Marks)
			case 3:
				student := GetStudentData()
				studentList.AddStudent(student)
			case 4:
				var id int
				fmt.Println("Enter the student's ID")
				fmt.Scanln(&id)
				studentList.RemoveStudentByID(id)
			case 5:
				student := GetStudentData()
				studentList.UpdateStudent(student)
			case 6:
				fmt.Println("you exited out of program!")
			default:
				fmt.Println("Invalid Input, try again")	
				
		}
	}
	data, err = json.Marshal(studentList)
	if err != nil {
		fmt.Println("Error in marshalling:", err)
		return
	}

	err = ioutil.WriteFile("../storage/jsons/student.json",data,0644)
	if err != nil {
		fmt.Println("Error in writing file:",err)
		return
	}

	
}

func GetStudentData() repo.Student {
	student := repo.Student{
        Section: "C",
        Subjects: []repo.Subject{
            {
                ID:   1,
                Name: "Maths",
            },
            {
                ID:   2,
                Name: "Science",
            },
            {
                ID:   3,
                Name: "English",
            },
        },
    }
	var name,faculty string
	var mathGrade,scienceGrade,englishGrade int

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the student's name:") 
	name, _ = reader.ReadString('\n')
	student.Name = name[:len(name)-1]
	fmt.Println("Enter the student's faculty's name:")
	faculty, _ = reader.ReadString('\n')
	student.Course = faculty[:len(faculty)-1]
	fmt.Println("Enter the student's math's grade")
	fmt.Scanln(&mathGrade)
	student.Subjects[0].Marks = mathGrade
	fmt.Println("Enter the student's science's grade")
	fmt.Scanln(&scienceGrade)
	student.Subjects[1].Marks = scienceGrade
	fmt.Println("Enter the student's english's grade")
	fmt.Scanln(&englishGrade)
	student.Subjects[2].Marks = englishGrade
	
	student.Grade = (student.Subjects[0].Marks + student.Subjects[1].Marks  + student.Subjects[2].Marks)/3

	return student
}