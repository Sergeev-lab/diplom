package main

import (

)

type address struct {
	Id string
	Name string
	Type type_
}

type commands_and_person struct {
	Id string
	Sorevnovanie sorevnovanie
	Commands commands
	Person person
	Number string
	Position string
}

type commands struct {
	Id string
	Name string
	Logo string
	Present string
	Sport sports
}

type user struct {
	Id string
	Login string
	Password string
	Command commands
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
	Fscore string
	Sscore string
	Status string
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

type rezults_command struct {
	Id string
	Command commands
	Sorev sorevnovanie
	Plase string
}

type rezults_sorev struct {
	Id string
	Sorevnovania sorevnovanie
	Command commands
	Points string
}

type sorevnovanie struct {
	Id string
	Name string
	Logo string
	Sport sports
	Fdata, Sdata string
	Level levels
	Country, Subject, City type_
	Stadium stadium
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

type stadium struct {
	Id string
	Name string
	Map string
}

// Пользовательские структуры //

type sorevnovanie_and_match struct {
	Sorevnovanie sorevnovanie
	Match []matches
}

type for_match_page struct {
	Admin bool
	Match matches
	Score string
	Fplayers []commands_and_person
	Splayers []commands_and_person
}

type for_commands_page struct {
	Info commands
	Dost []rezults_command
	Results []matches
	Calendar []matches
}

type for_sorevnovanie_page struct {
	Sorevnovanie sorevnovanie
	Commands []commands
	History []matches
	Live []matches
	Kalendar []matches
	Points []rezults_sorev
}

type for_user_page struct {
	Data user
	Dost []rezults_command
	Sorev []sorevnovanie
}