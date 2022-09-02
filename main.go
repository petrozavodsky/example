package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"students/pkg/storage"
	"students/pkg/student"
)

func main() {

	// инициализируем хранилище
	newStorage := storage.NewStudentsStorage()

	fmt.Println("Введите имя, возраст и курс через пробел и нажмите  enter:")

	var in = bufio.NewReader(os.Stdin)

	for {

		line, err := in.ReadString('\n')

		if err == io.EOF {
			fmt.Println("Все студенты:")
			break
		}

		row := strings.Fields(line)

		if len(row) < 3 {
			fmt.Print("Необходимо ввести данные студента\n")
			continue
		}

		name := row[0]

		age, err := strconv.Atoi(row[1])

		if err != nil {
			fmt.Print("Ошибка ввода возраста\n")
			continue
		}

		grade, err := strconv.Atoi(row[2])

		if err != nil {
			fmt.Print("Ошибка ввода возраста\n")
			continue
		}

		// создаем одного ноывого сутдента
		studentSingle := student.NewStudent(name, age, grade)

		// добавляем пользователя в хранилище
		newStorage.Put(studentSingle)
	}

	// перебираем всех пользователей из хранилища при помощи метода GetAll на мой взгляд он нужен и хорош
	for _, value := range newStorage.GetAll() {

		// создаем переменную чтолько для того что бы сущетвование метода Get как то опрадать в этом примере кода
		user, _ := newStorage.Get(value.Name)

		fmt.Println(user.Name, user.Age, user.Grade)
	}
}
