// main.go
package main

import (
	"virtbro/cmd"
	"virtbro/pkg/db"
)

func main() {
	db.InitDB()
	defer db.CloseDB()
	cmd.Execute()
}
