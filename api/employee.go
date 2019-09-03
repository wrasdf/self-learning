package employee

type Employee interface {
	Create()
	Update()
}


type Employee struct {
  Id          int
  Name        string
  Department  string
  Email       string
  Phone       string
}
