package main

import (

)

type address struct {
	Id string
	Name string
	Type type_
}

type commands struct {
	Id string
	Name string
	Logo string
	Present string
	Sport sports
}

type commands_and_person struct {
	Id string
	Commands commands
	Person person
	Number string
}

type levels struct {
	Id string
	Name string
}

type matches struct {
	Id string
	Fcommand commands
	Scommand commands
	Data string
	Sorevnovanie sorevnovanie
	Total string
}

type person struct {
	Id string
	Fio string
	Logo string
	Type type_
}

type slider struct {
	Id string
	Img string
	Title string
	Subtitle string
	Description string
	Btn string
}

type sorevnovanie struct {
	Id string
	Name string
	Logo string
	Sport sports
	Fdata, Sdata string
	Level levels
	Country, Subject, City type_
}

type sorevnovania_and_commands struct {
	Id string
	Sorevnovanie sorevnovanie
	Command commands
}

type sports struct {
	Id string
	Name string
	Logo string
	Href string
}

type type_ struct {
	Id string
	Name string
}