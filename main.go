package main

import (
	"fmt"
	"userInfo/models"
)

func main() {
	fmt.Println("Hello World")
	models.DbConnection()
	models.Dbtst()
}
