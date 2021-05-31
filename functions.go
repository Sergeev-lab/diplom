package main

import (
	"fmt"
	"time"
	"errors"
	jwt "github.com/dgrijalva/jwt-go"
)

type result struct {
	Match_id string
	Fc_id string
	Fc_name string
	Fc_logo string
	Fc_present string
	Sc_id string
	Sc_name string
	Sc_logo string
	Sc_present string
	Total string
	Sorev_id string
	Sorev_name string
	Data string
}

type data struct {
	Info commands
	Results []result
	Kalendar []result
}

type tablepoints struct {
	Position int
	Id string
	Name string
	Logo string
	Present string
	Points string
}

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
	res, err := database.Query("SELECT DISTINCT sorev.id, sorev.name FROM `matches` JOIN sorevnovania AS sorev ON sorev.id = matches.sorevnovania_id WHERE matches.data < ? AND matches.status = 'live' AND sorev.sport_id = ?", data, id)
	if err != nil {
		fmt.Println(err)
	}
	for res.Next() {
		p := sorev {}
		res.Scan(&p.Id, &p.Name)
		rest, _ := database.Query("SELECT matches.id, Fc.name, Sc.name, matches.total FROM matches JOIN commands as Fc ON Fc.id = matches.fcommand_id JOIN commands as Sc ON Sc.id = matches.scommand_id WHERE matches.sorevnovania_id = ? AND matches.status = 'live'", p.Id)
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

func Match(id string) formatch {
	p := formatch {}
	name := []player {}

	// Информация о матче
	res, err := database.Query("SELECT fc.id, fc.name, fc.present, fc.logo, sc.id, sc.name, sc.present, sc.logo, sorev.id, sorev.name, total, city.name, stad.name, data FROM matches JOIN commands AS fc ON fc.id = matches.fcommand_id JOIN commands AS sc ON sc.id = matches.scommand_id JOIN sorevnovania AS sorev ON sorev.id = matches.sorevnovania_id JOIN address AS city ON city.id = sorev.city_id JOIN address AS stad ON stad.id = sorev.stadium_id WHERE matches.id = ?", id)
	if err != nil {
		fmt.Println(err)
	}
	for res.Next() {
		res.Scan(&p.Fc_id, &p.Fc_name, &p.Fc_present, &p.Fc_logo, &p.Sc_id, &p.Sc_name, &p.Sc_present, &p.Sc_logo, &p.Sorev_id, &p.Sorev_name, &p.Total, &p.City, &p.Stad, &p.Data)
		t, _ := time.Parse("2006-01-02 15:04:05", p.Data)
		p.Data = t.Format("2 January 2006 15:04")
	}	
	
	// Список игроков 1ой команды
	res2, err := database.Query("SELECT number, person.fio, position FROM `commands_and_person` JOIN person ON person.id = person_id WHERE sorevnovania_id = ? AND commands_id = ?", p.Sorev_id, p.Fc_id)
	if err != nil {
		fmt.Println(err)
	}
	for res2.Next() {
		s := player {}
		res2.Scan(&s.Number, &s.Name, &s.Position)
		name = append(name, s)
	}
	p.Fc_players = name
	name = nil

	// Список игроков 2ой команды
	res3, err := database.Query("SELECT number, person.fio, position FROM `commands_and_person` JOIN person ON person.id = person_id WHERE sorevnovania_id = ? AND commands_id = ?", p.Sorev_id, p.Sc_id)
	if err != nil {
		fmt.Println(err)
	}
	for res3.Next() {
		ss := player {}
		res3.Scan(&ss.Number, &ss.Name, &ss.Position)
		name = append(name, ss)
	}
	p.Sc_players = name
	name = nil

	return p
}

func Commands(id string) data {
	name := data {}

	// Информация о команде
	res1, err := database.Query("SELECT commands.id, commands.name, commands.logo, commands.present, sports.name FROM `commands` JOIN sports ON sports.id = sports_id WHERE commands.id = ?", id)
	if err != nil {
		fmt.Println(err)
	}
	for res1.Next() {
		res1.Scan(&name.Info.Id, &name.Info.Name, &name.Info.Logo, &name.Info.Present, &name.Info.Sport.Name)
	}

	// Результаты команды
	a := []result {}
	res2, err := database.Query("SELECT matches.id, fc.id, fc.name, fc.logo, fc.present, sc.id, sc.name, sc.logo, sc.present, matches.total, sorev.id, sorev.name, matches.data FROM matches JOIN commands AS fc ON fc.id = matches.fcommand_id JOIN commands AS sc ON sc.id = matches.scommand_id JOIN sorevnovania AS sorev ON sorev.id = matches.sorevnovania_id WHERE matches.status = 'finish' AND (fc.id = ? OR sc.id = ?)", id, id)
	if err != nil {
		fmt.Println(err)
	}
	for res2.Next() {
		b := result {}
		res2.Scan(&b.Match_id, &b.Fc_id, &b.Fc_name, &b.Fc_logo, &b.Fc_present, &b.Sc_id, &b.Sc_name, &b.Sc_logo, &b.Sc_present, &b.Total, &b.Sorev_id, &b.Sorev_name, &b.Data)
		t, _ := time.Parse("2006-01-02 15:04:05", b.Data)
		b.Data = t.Format("2 January 2006 15:04")
		a = append(a, b)
	}
	name.Results = a

	// Календарь команды
	aa := []result {}
	res3, err := database.Query("SELECT matches.id, fc.id, fc.name, fc.logo, fc.present, sc.id, sc.name, sc.logo, sc.present, sorev.id, sorev.name, matches.data FROM matches JOIN commands AS fc ON fc.id = matches.fcommand_id JOIN commands AS sc ON sc.id = matches.scommand_id JOIN sorevnovania AS sorev ON sorev.id = matches.sorevnovania_id WHERE matches.status = 'up_coming' AND (fc.id = ? OR sc.id = ?)", id, id)
	if err != nil {
		fmt.Println(err)
	}
	for res3.Next() {
		bb := result {}
		res3.Scan(&bb.Match_id, &bb.Fc_id, &bb.Fc_name, &bb.Fc_logo, &bb.Fc_present, &bb.Sc_id, &bb.Sc_name, &bb.Sc_logo, &bb.Sc_present, &bb.Sorev_id, &bb.Sorev_name, &bb.Data)
		t, _ := time.Parse("2006-01-02 15:04:05", bb.Data)
		bb.Data = t.Format("2 January 2006 15:04")
		aa = append(aa, bb)
	}
	name.Kalendar = aa

	return name
}

func Sorevnivania(id string) []tablepoints {
	// Таблица очков
	name := []tablepoints {}
	i := 1
	res1, err := database.Query("SELECT commands.id, commands.name, commands.logo, commands.present, sorevnovania_and_commands.points FROM sorevnovania_and_commands JOIN commands ON commands.id = sorevnovania_and_commands.commands_id WHERE sorevnovania_and_commands.sorevnovania_id = ? ORDER BY sorevnovania_and_commands.points DESC", id)
	if err != nil {
		fmt.Println(err)
	}
	for res1.Next() {
		p := tablepoints {}
		res1.Scan(&p.Id, &p.Name, &p.Logo, &p.Present, &p.Points)
		p.Position = i
		name = append(name, p)
		i++
	}
	return name
}

func getSorevName(id string) string {
	var name string
	res1 := database.QueryRow("SELECT name FROM `sorevnovania` WHERE id = ?", id)
	res1.Scan(&name)

	return name
}

func getToken(mySigningKey []byte) string {
	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, _ := token.SignedString(mySigningKey)

	return tokenString
}

func parseToken(tokenString string, mySigningKey []byte) bool {
	token, err :=jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(mySigningKey), nil
	})

	if err == nil && token.Valid {
        // fmt.Println("Your token is valid.  I like your style.")
		return true
    } else {
        // fmt.Println("This token is terrible!  I cannot accept this.")
		return false
    }
}

func register(u, p string) error {
	name := user {}
	res := database.QueryRow("SELECT * FROM `users` WHERE login = ?", u)
	res.Scan(&name.Id, &name.Login, &name.Password, &name.Type_id, &name.Token)

	if len(name.Id) > 0 {
		return errors.New("Такой пользователь уже существует")
	} else {
		token := getToken(mySigningKey)
		_, err := database.Exec("INSERT INTO `users` (`id`, `login`, `password`, `type_id`, `token`) VALUES (NULL, ?, ?, '9', ?)", u, p, token)
		if err != nil {
			return err
		}
	}

	return nil
}

func login(l, p string) error {
	name := user {}
	res := database.QueryRow("SELECT * FROM `users` WHERE login = ?", l)
	res.Scan(&name.Id, &name.Login, &name.Password, &name.Type_id, &name.Token)

	if len(name.Id) == 0 {
		return errors.New("Пользователь не найден")
	}
	if name.Password != p {
		return errors.New("Неверный пароль")
	}
	if parseToken(name.Token, mySigningKey) == true {
		fmt.Println("Токен валидный!")
	} else {
		return errors.New("Токен не валидный")
	}

	return nil
}