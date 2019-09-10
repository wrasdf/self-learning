package employees

import (
	"fmt"
	"models/employees"
	"models/departments"
	"github.com/mitchellh/mapstructure"
)

type Employees interface {
	init()
	all() 													[]Employee
	get(id string) 									Employee
	add(obj Employee) 							Employee
	// patch(obj Employee) 						Employee
	// update(obj map[string]string) 	Employee
	delete(id string) 							Employee
}

type Employees struct {
	Employees []Employee
}

type Department struct {
	Id         		string
	Name       		string
	Description   string
}

type Employee struct {
	Id         string
	Name       string
	Email      string
	Phone      string
	Department Department
}

func NewEmployees() {
	workers = employees.all()
	for i := 0; i < workers.len; i++ {
		Decode(input, &result)
	}
	c.Employees = employees.all() // ?
}

func (c *Employees) all() []Employee {
	// need transform ?
	return employees.all()
}

func (c *Employees) get(id string) Employee {
	employee = employees.get(employee)
	return employee
}

func (c *Employees) add(employee Employee) Employee {
	employees.add(employee)
	c.Employees = append(c.Employees, employee)
	return employeed
}

func (c *Employees) delete(id string) Employee {
	employee := models.employees.delete(id)
	c.Employees.remove(employee)
	return employee
}
