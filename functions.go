package main

import (
	"fmt"
)

func Slider() []slider {
	name := []slider {} 

	res, err := database.Query("SELECT * FROM `slider`")
	if err != nil {
		fmt.Println(err)
	}

	for res.Next(){
		p := slider {}
		res.Scan(&p.Id, &p.Img, &p.Title, &p.Subtitle, &p.Description, &p.Btn)
		name = append(name, p)
	}

	return name
}