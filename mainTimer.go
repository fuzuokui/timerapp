package main

import (
	"fmt"
	"timerapp/models"
)

func main() {
	var bar models.Bar
	var goBar models.GoBar
	data, err := bar.QueryAddBars(4)
	if err != nil {
		fmt.Println(err)
	} else {
		goBar.Add(data)
	}

}
