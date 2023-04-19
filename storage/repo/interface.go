package repo


type Subject struct {
    ID   int `json:"id"`
    Name string `json:"name"`
	Marks int `json:"marks"`
}

type StudentManager interface{
	GetStudents()[]Student
	GetStudentByID(id int)(Student)
	AddStudent(student Student)
	RemoveStudentByID(id int)
	UpdateStudent(s Student)

}

type Student struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Grade    int    `json:"grade"`
	Section  string `json:"section"`
	Course   string `json:"course"`
	Subjects []Subject `json:"subjects"`
}

type StudentList struct {
	Students []Student
}