package main

import (
	"fmt"
	"time"
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

func Sorev(id int) []sorev {
	data := time.Now().Format("2006-01-02 15:04")
	sor := []sorev {}
	mat := []math {}
	res, err := database.Query("SELECT DISTINCT sorev.id, sorev.name FROM `matches` JOIN sorevnovania AS sorev ON sorev.id = matches.sorevnovania_id WHERE matches.data < ? AND matches.status IS null AND sorev.sport_id = ?", data, id)
	if err != nil {
		fmt.Println(err)
	}
	for res.Next() {
		p := sorev {}
		res.Scan(&p.Id, &p.Name)
		rest, _ := database.Query("SELECT matches.id, Fc.name, Sc.name, matches.total FROM matches JOIN commands as Fc ON Fc.id = matches.fcommand_id JOIN commands as Sc ON Sc.id = matches.scommand_id WHERE matches.sorevnovania_id = ? AND matches.status IS null", p.Id)
		for rest.Next() {
			a:= math {}
			rest.Scan(&a.Id, &a.Fc, &a.Sc, &a.Total)
			mat = append(mat, a)
		}
		p.Match = mat
		mat = nil
		sor = append(sor, p)
	}
	return sor
}