package jsons

import(
	"Student_Management/storage/repo"
)

type Students repo.StudentList

func (s *Students) GetStudents() []repo.Student {
	return s.Students
}

func (s *Students) AddStudent(student repo.Student) {
	student.ID = len(s.Students) + 1
	student.Grade = func() int {
		var sum int
		for _,val := range student.Subjects {
			sum +=  val.Marks
		}
		return sum/len(student.Subjects)
	}()
	s.Students = append(s.Students, student)
}

func (s *Students) GetStudentByID(id int) (repo.Student){
	var student repo.Student
	for i, val := range s.Students {
		if id == val.ID {
			student = s.Students[i]
			break
		}
	}
	return student
}

func (s *Students) RemoveStudentByID(id int){
	for i := 0; i < len(s.Students); i++ {
		if id == s.Students[i].ID {
			s.Students = append(s.Students[:i], s.Students[i+1:]...)
		}
	}
}

func (s *Students) UpdateStudent(student repo.Student){
	for i, val := range s.Students {
		if student.Name == val.Name {
			student.ID = val.ID
			s.Students[i] = student
		}
	}
}