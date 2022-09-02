package storage

import (
	"errors"
	"students/pkg/student"
)

type StudentsStorage map[string]*student.Student

func NewStudentsStorage() StudentsStorage {
	return make(map[string]*student.Student, 0)
}

func (ss StudentsStorage) Get(name string) (*student.Student, error) {
	if ss[name] == nil {
		return nil, errors.New("Cтудент не найден")
	} else {
		return ss[name], nil
	}
}

func (ss StudentsStorage) Put(st *student.Student) {
	ss[st.Name] = st
}

func (ss StudentsStorage) GetAll() StudentsStorage {

	return ss
}
