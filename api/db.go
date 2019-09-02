package db

import (
  "fmt",
  "database/sql",
  _ "github.com/go-sql-driver/mysql"
)

func NewDBController() {

}

type DBController interface {
	Connect()
	Query(query string)
	Update(query string)
}

type DBController struct {
  dbDriver    string
  dbUser      string
  dbPassWord  string
  dbName      string
}
