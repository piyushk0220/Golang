package main

import (
	"fmt"
)

type Employee struct {
	id     int
	name   string
	salary int
	skill  string
}

func main() {

	employees := make(map[int]Employee)

	var input int

	for {

		fmt.Println("Welcome to Employee Management App")
		fmt.Println("Press 1. Create")
		fmt.Println("Press 2. Read")
		fmt.Println("Press 3. Update")
		fmt.Println("Press 4. Delete")

		fmt.Scanln(&input)

		if input == 1 {
			createEmp(employees)
		} else if input == 2 {
			readEmp(employees)
		} else if input == 3 {
			updateEmp(employees)
		} else if input == 4 {
			deleteEmp(employees)
		} else {
			break
		}

	}

}

func createEmp(employees map[int]Employee) {
	emp := Employee{}
	fmt.Println("Enter employee details")
	fmt.Println("Enter employee id")
	fmt.Scanln(&emp.id)
	fmt.Scanln(&emp.name)
	fmt.Scanln(&emp.skill)
	fmt.Scanln(&emp.salary)
	//fmt.Printf("%#v \n", emp)
	employees[emp.id] = emp

}
func updateEmp(employees map[int]Employee) {

	var id, sal int

	fmt.Println("Enter id whose salary wanna icrease")
	fmt.Scanln(&id)
	fmt.Println("Enter new salary")
	fmt.Scanln(&sal)

	updatedEmp, ok := employees[id]
	if ok == true {
		updatedEmp.salary = sal
		employees[id] = updatedEmp
	} else {
		fmt.Println("id doesnt exist")
	}

}
func readEmp(employees map[int]Employee) {
	fmt.Printf("%#v \n", employees)

}
func deleteEmp(employees map[int]Employee) {
	var id int
	fmt.Println("Enter id u wanna delete")
	fmt.Scanln(&id)
	delete(employees, id)

}
