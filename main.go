package main

import (
	_ "database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var a int = 6
	fmt.Printf("%v", a*a)
}
