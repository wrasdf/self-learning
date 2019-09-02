package employee

type Employee interface {
	Init()
	Update()
  Save()
}


type Employee struct {
  Id          int
  Name        string
  Department  string
  Email       string
  Phone       string
}
