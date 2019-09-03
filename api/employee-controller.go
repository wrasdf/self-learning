package econtroller

import (
  "employee"
  "db"
)

type EController interface {
	all() List(employee)
	get() employee
  add() employee
  update() employee
  delete() employee
}


type Employee struct {

}

func ()
