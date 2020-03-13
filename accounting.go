package main

import (
	"fmt"
	"golandAccounting/domain/model"
)

func main() {
	var item = model.Item{
		Name:        "Table",
		Description: "table of eight chairs ",
		Price:       1200,
	}

	fmt.Println(item)
}
