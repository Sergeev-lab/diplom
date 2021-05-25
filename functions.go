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

func Match() []matches {
	name := []matches {}

	res, err := database.Query("SELECT Fcommand.name, Scommand.name, sorev.name, sports.name, sports.logo, sports.href FROM matches JOIN commands AS Fcommand ON Fcommand.id = matches.fcommand_id JOIN commands AS Scommand ON Scommand.id = matches.scommand_id JOIN sorevnovania AS sorev ON sorev.id = matches.sorevnovania_id JOIN sports ON sports.id = sorev.sport_id")
	if err != nil {
		fmt.Println(err)
	}

	for res.Next(){
		p := matches {}
		res.Scan(&p.Fcommand.Name, &p.Scommand.Name, &p.Sorevnovanie.Name, &p.Sorevnovanie.Sport.Name, &p.Sorevnovanie.Sport.Logo, &p.Sorevnovanie.Sport.Href)
		name = append(name, p)
	}

	return name
	// Field_hockey: matches {		
	// 	Fcommand: commands {
	// 		Name: "command_name",
	// 	},
	// 	Scommand: commands {
	// 		Name: "command_name",
	// 	},
	// 	Sorevnovanie: sorevnovanie {
	// 		Name: "sorevnovanie_name",
	// 		Sport: sports {
	// 			Name: "sports_name",
	// 			Logo: "sports_logo",
	// 			Href: "/href/",
	// 		},
	// 	},
	// },
}