package departments

type Departments interface {
	all()
	get()
	add()
	patch()
	update()
	delete()
}


type Departments struct {
  Id           string
  Name         string
  Description  string
}
