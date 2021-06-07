package main

import (
	"fmt"
	"time"
	"errors"
	jwt "github.com/dgrijalva/jwt-go"
	"net/url"
	"mime/multipart"
	"io/ioutil"
	"strings"
	"net/http"
)

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

func Sorev(id int) []sorevnovanie_and_match {
	data := time.Now().Format("2006-01-02 15:04")
	sor := []sorevnovanie_and_match {}
	mat := []matches {}

	res, err := database.Query("SELECT DISTINCT sorev.id, sorev.name FROM `matches` JOIN sorevnovania AS sorev ON sorev.id = matches.sorevnovania_id WHERE matches.data < ? AND matches.status = 'live' AND sorev.sport_id = ?", data, id)
	if err != nil {
		fmt.Println(err)
	}
	for res.Next() {
		p := sorevnovanie_and_match {}
		res.Scan(&p.Sorevnovanie.Id, &p.Sorevnovanie.Name)
		rest, _ := database.Query("SELECT matches.id, Fc.name, Sc.name, matches.fscore, matches.sscore FROM matches JOIN commands_or_players as Fc ON Fc.id = matches.fcommand_id JOIN commands_or_players as Sc ON Sc.id = matches.scommand_id WHERE matches.sorevnovania_id = ? AND matches.status = 'live'", p.Sorevnovanie.Id)
		for rest.Next() {
			a := matches {}
			rest.Scan(&a.Id, &a.Fcommand.Name, &a.Scommand.Name, &a.Fscore, &a.Sscore)
			mat = append(mat, a)
		}
		p.Match = mat
		mat = nil
		sor = append(sor, p)
	}

	return sor
}

func Match(id string) for_match_page {
	p := for_match_page {}
	f1 := []commands_and_person {}
	f2 := []commands_and_person {}

	// Информация о матче
	res, err := database.Query("SELECT fc.id, fc.name, fc.present, fc.logo, sc.id, sc.name, sc.present, sc.logo, sorev.id, sorev.name, fscore, sscore, city.name, stad.name, data FROM matches JOIN commands_or_players AS fc ON fc.id = matches.fcommand_id JOIN commands_or_players AS sc ON sc.id = matches.scommand_id JOIN sorevnovania AS sorev ON sorev.id = matches.sorevnovania_id JOIN address AS city ON city.id = sorev.city_id JOIN address AS stad ON stad.id = sorev.stadium_id WHERE matches.id = ?", id)
	if err != nil {
		fmt.Println(err)
	}
	for res.Next() {
		res.Scan(&p.Match.Fcommand.Id, &p.Match.Fcommand.Name, &p.Match.Fcommand.Present, &p.Match.Fcommand.Logo, &p.Match.Scommand.Id, &p.Match.Scommand.Name, &p.Match.Scommand.Present, &p.Match.Scommand.Logo, &p.Match.Sorevnovanie.Id, &p.Match.Sorevnovanie.Name, &p.Match.Fscore, &p.Match.Sscore, &p.Match.Sorevnovanie.City.Name, &p.Match.Sorevnovanie.Stadium.Name, &p.Match.Data)
		t, _ := time.Parse("2006-01-02 15:04:05", p.Match.Data)
		p.Match.Data = t.Format("2 January 2006 15:04")
	}	
	
	// Список игроков 1ой команды
	res2, err := database.Query("SELECT number, person.fio, position FROM `commands_and_person` JOIN person ON person.id = person_id WHERE sorevnovania_id = ? AND commands_id = ?", p.Match.Sorevnovanie.Id, p.Match.Fcommand.Id)
	if err != nil {
		fmt.Println(err)
	}
	for res2.Next() {
		s := commands_and_person {}
		res2.Scan(&s.Number, &s.Person.Fio, &s.Position)
		f1 = append(f1, s)
	}
	p.Fplayers = f1

	// Список игроков 2ой команды
	res3, err := database.Query("SELECT number, person.fio, position FROM `commands_and_person` JOIN person ON person.id = person_id WHERE sorevnovania_id = ? AND commands_id = ?", p.Match.Sorevnovanie.Id, p.Match.Scommand.Id)
	if err != nil {
		fmt.Println(err)
	}
	for res3.Next() {
		ss := commands_and_person {}
		res3.Scan(&ss.Number, &ss.Person.Fio, &ss.Position)
		f2 = append(f2, ss)
	}
	p.Splayers = f2
	return p
}

func Commands(id string) for_commands_page {
	name := for_commands_page {}

	// Информация о команде
	res1 := database.QueryRow("SELECT commands_or_players.id, commands_or_players.name, commands_or_players.logo, commands_or_players.present, sports.name FROM `commands_or_players` JOIN sports ON sports.id = sports_id WHERE commands_or_players.id = ?", id)
	res1.Scan(&name.Info.Id, &name.Info.Name, &name.Info.Logo, &name.Info.Present, &name.Info.Sport.Name)
	
	// Получаем информацию о достижениях
	dd := []rezults_command {}
	d := rezults_command {}
	res, err := database.Query("SELECT rezults_command.plase, rezults_command.sorevnovanie_id, sorev.name FROM `rezults_command` JOIN sorevnovania AS sorev ON sorev.id = rezults_command.sorevnovanie_id WHERE commands_or_players_id = ?", name.Info.Id)
	if err != nil {
		fmt.Println(err)
	}
	for res.Next() {
		res.Scan(&d.Plase, &d.Sorev.Id, &d.Sorev.Name)
		dd = append(dd, d)
	}
	name.Dost = dd
	
	// Результаты команды
	a := []matches {}
	res2, err := database.Query("SELECT matches.id, fc.id, fc.name, fc.logo, fc.present, sc.id, sc.name, sc.logo, sc.present, matches.fscore, matches.sscore, sorev.id, sorev.name, matches.data FROM matches JOIN commands_or_players AS fc ON fc.id = matches.fcommand_id JOIN commands_or_players AS sc ON sc.id = matches.scommand_id JOIN sorevnovania AS sorev ON sorev.id = matches.sorevnovania_id WHERE matches.status = 'finish' AND (fc.id = ? OR sc.id = ?)", id, id)
	if err != nil {
		fmt.Println(err)
	}
	for res2.Next() {
		b := matches {}
		res2.Scan(&b.Id, &b.Fcommand.Id, &b.Fcommand.Name, &b.Fcommand.Logo, &b.Fcommand.Present, &b.Scommand.Id, &b.Scommand.Name, &b.Scommand.Logo, &b.Scommand.Present, &b.Fscore, &b.Sscore, &b.Sorevnovanie.Id, &b.Sorevnovanie.Name, &b.Data)
		t, _ := time.Parse("2006-01-02 15:04:05", b.Data)
		b.Data = t.Format("2 January 2006 15:04")
		a = append(a, b)
	}
	name.Results = a

	// Календарь команды
	aa := []matches {}
	res3, err := database.Query("SELECT matches.id, fc.id, fc.name, fc.logo, fc.present, sc.id, sc.name, sc.logo, sc.present, sorev.id, sorev.name, matches.data FROM matches JOIN commands_or_players AS fc ON fc.id = matches.fcommand_id JOIN commands_or_players AS sc ON sc.id = matches.scommand_id JOIN sorevnovania AS sorev ON sorev.id = matches.sorevnovania_id WHERE matches.status = 'up_coming' AND (fc.id = ? OR sc.id = ?)", id, id)
	if err != nil {
		fmt.Println(err)
	}
	for res3.Next() {
		bb := matches {}
		res3.Scan(&bb.Id, &bb.Fcommand.Id, &bb.Fcommand.Name, &bb.Fcommand.Logo, &bb.Fcommand.Present, &bb.Scommand.Id, &bb.Scommand.Name, &bb.Scommand.Logo, &bb.Scommand.Present, &bb.Sorevnovanie.Id, &bb.Sorevnovanie.Name, &bb.Data)
		t, _ := time.Parse("2006-01-02 15:04:05", bb.Data)
		bb.Data = t.Format("2 January 2006 15:04")
		aa = append(aa, bb)
	}
	name.Calendar = aa

	return name
}

func Sorevnivania(id string) for_sorevnovanie_page {
	s := for_sorevnovanie_page {}

	// Информация о соревновании
	p := sorevnovanie {}
	res1 := database.QueryRow("SELECT sorevnovania.name, sorevnovania.logo, sorevnovania.fdata, sorevnovania.sdata, levels.name, country.name, subject.name, city.name, stadium.name, sorevnovania.map FROM `sorevnovania` JOIN levels ON levels.id = sorevnovania.level_id JOIN address AS country ON country.id = sorevnovania.country_id JOIN address AS subject ON subject.id = sorevnovania.subject_id JOIN address AS city ON city.id = sorevnovania.city_id JOIN address AS stadium ON stadium.id = sorevnovania.stadium_id WHERE sorevnovania.id = ?", id)
	res1.Scan(&p.Name, &p.Logo, &p.Fdata, &p.Sdata, &p.Level.Name, &p.Country.Name, &p.Subject.Name, &p.City.Name, &p.Stadium.Name, &p.Map)
	s.Sorevnovanie = p

	// Информация о участниках соревнорвания
	cc := []commands {}
	c := commands {}
	res2, err := database.Query("SELECT command.id, command.name, command.logo, command.present FROM `sorevnovania_and_commands` JOIN commands_or_players AS command ON command.id = sorevnovania_and_commands.commands_id WHERE sorevnovania_and_commands.sorevnovania_id = ?", id)
	if err != nil {
		fmt.Println(err)
	}
	for res2.Next() {
		res2.Scan(&c.Id, &c.Name, &c.Logo, &c.Present)
		cc = append(cc, c)
	}
	s.Commands = cc

	// Информация о таблице очков
	data := time.Now().Format("2006-01-02")
	tt := []rezults_sorev {}
	if p.Fdata < data || p.Fdata == data {
		t := rezults_sorev {}
		res3, err := database.Query("SELECT command.id, command.name, command.logo, rezults_sorev.points FROM `rezults_sorev` JOIN commands_or_players AS command ON command.id = rezults_sorev.commands_or_players_id WHERE rezults_sorev.sorevnovania_id = ? ORDER BY `rezults_sorev`.`points` DESC", id)
		if err != nil {
			fmt.Println(err)
		}
		for res3.Next() {
			res3.Scan(&t.Command.Id, &t.Command.Name, &t.Command.Logo, &t.Points)
			tt = append(tt, t)
		}
		s.Points = tt
	}

	return s
}

func getSorev(id string) sorevnovanie {
	var s sorevnovanie

	res1 := database.QueryRow("SELECT sorevnovania.id, sorevnovania.name, sorevnovania.logo, sorevnovania.fdata, sorevnovania.sdata, lvl.name, country.name, subj.name, city.name, stad.name, sorevnovania.map FROM `sorevnovania` JOIN levels AS lvl ON lvl.id = sorevnovania.level_id  JOIN address AS country ON country.id = sorevnovania.country_id JOIN address AS subj ON subj.id = sorevnovania.subject_id JOIN address AS city ON city.id = sorevnovania.city_id JOIN address AS stad ON stad.id = sorevnovania.stadium_id WHERE sorevnovania.id = ?", id)
	res1.Scan(&s.Id, &s.Name, &s.Logo, &s.Fdata, &s.Sdata, &s.Level.Name, &s.Country.Name, &s.Subject.Name, &s.City.Name, &s.Stadium.Name, &s.Map)
	
	t1, _ := time.Parse("2006-01-02", s.Fdata)
	t2, _ := time.Parse("2006-01-02", s.Sdata)
	s.Fdata = t1.Format("2 January 2006")
	s.Sdata = t2.Format("2 January 2006")

	return s
}

func getPlayers(id string) []commands {
	com := []commands {}
	res, _ := database.Query("SELECT commands.id, commands.name, commands.logo, commands.present FROM `sorevnovania_and_commands` JOIN commands_or_players AS commands ON commands.id = sorevnovania_and_commands.commands_id WHERE sorevnovania_and_commands.sorevnovania_id = ?", id)
	for res.Next() {
		p := commands {}
		res.Scan(&p.Id, &p.Name, &p.Logo, &p.Present)
		com = append(com, p)
	}

	return com
}

func getToken(mySigningKey []byte, id int8) string {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	claims["User_id"] = id
	token.Claims = claims

	tokenString, _ := token.SignedString(mySigningKey)

	return tokenString
}

func parseToken(tokenString string, mySigningKey []byte) (jwt.MapClaims, bool) {
	token, err :=jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(mySigningKey), nil
	})

	if err == nil && token.Valid {
        claims := token.Claims.(jwt.MapClaims)
		return claims, true
    } else {
		return nil, false
    }
}

func checkUser(u, p string) (string, error) {
	name := user {}
	res := database.QueryRow("SELECT * FROM `users` WHERE login = ?", u)
	res.Scan(&name.Id, &name.Login, &name.Password, &name.Command.Id)

	if len(name.Id) > 0 {
		return "", errors.New("Такой пользователь уже существует")
	} else {
		res, err := database.Exec("INSERT INTO `users` (`id`, `login`, `password`, `type_id`) VALUES (NULL, ?, ?, '9')", u, p)
		if err != nil {
			return "", err
		}
		id, _ := res.LastInsertId()
		
		token := getToken(mySigningKey, int8(id))

		return token, nil
	}
}

func register(w http.ResponseWriter, r *http.Request) (error) {
	r.ParseMultipartForm(32 << 20)

	// Проверяем есть ли такой пользователь с таким паролем в БД
	token, err := checkUser(r.Form.Get("login"), r.Form.Get("password"))
	if err != nil {
		return err
	}

	// Скачиваем изображение
	pic, _, err := r.FormFile("Logo")
	if err != nil {
		return err
	}
	file, err := download(pic, "img/commands/")
	if err != nil {
		return err
	}

	// Создаем новую команду
	newCommand_id, er := newCommand(r.Form, file)
	if er != nil {
		return err
	}

	// Создаем нового пользователя
	e := newUser(r.Form.Get("login"), r.Form.Get("password"), newCommand_id)
	if e != nil {
		return err
	}

	// Устанавливаем токен в куки
	cookie := http.Cookie{Name: "token", Path: "/", Value: token}
	http.SetCookie(w, &cookie)

	// Выполняем переадресацию
	http.Redirect(w, r, "/", http.StatusSeeOther)
	
	return nil
}

func login(n, p string) (int8, error) {
	var id int8
	row := database.QueryRow("SELECT id FROM `users` WHERE users.login=? AND users.password=?", n, p)
	row.Scan(&id)
	if id == 0 {
		return id, errors.New("Не верный логин или пароль")
	}

	return id, nil
}

func cookieToken(w http.ResponseWriter, r *http.Request) (string, error) {
	cookie, err := r.Cookie("token")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return "", err
	}

	return cookie.Value, nil
}

func download(pic multipart.File, path string) (string, error) {
	fileBytes, err := ioutil.ReadAll(pic)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	tempFile, err := ioutil.TempFile(path, "upload-*.png")
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	tempFile.Write(fileBytes)

	file := tempFile.Name()
	Fname := strings.ReplaceAll(file, "\\", "/")

	return Fname, nil
}

func newCommand(f url.Values, logo string) (string, error) {
	res, err := database.Exec("INSERT INTO `commands` (`id`, `name`, `logo`, `present`, `sports_id`) VALUES (NULL, ?, ?, ?, ?)", f.Get("name"), logo, f.Get("city"), f.Get("sport"))
	if err != nil {
		return "", err
	}

	id, _ := res.LastInsertId()

	return string(id), nil
}

func newUser(l, p, id string) error {
	_, err := database.Exec("INSERT INTO `users` (`id`, `login`, `password`, `command_or_player_id`) VALUES (NULL, ?, ?, ?)", l, p, id)
	if err != nil {
		return err
	}
	
	return nil
}

func calendar(id string) []sorevnovanie {
	send := []sorevnovanie {}
	data := time.Now().Format("2006-01-02 15:04")
	res, err := database.Query("SELECT sorevnovania.id, sorevnovania.name, sorevnovania.fdata, sorevnovania.sdata, lvl.name, country.name, subj.name, city.name, stad.name FROM `sorevnovania` JOIN levels AS lvl ON lvl.id = sorevnovania.level_id  JOIN address AS country ON country.id = sorevnovania.country_id JOIN address AS subj ON subj.id = sorevnovania.subject_id JOIN address AS city ON city.id = sorevnovania.city_id JOIN address AS stad ON stad.id = sorevnovania.stadium_id WHERE  sorevnovania.fdata > ? AND sport_id = ?", data, id)
	if err != nil {
		fmt.Println(err)
	}
	for res.Next() {
		sor := sorevnovanie {}
		res.Scan(&sor.Id, &sor.Name, &sor.Fdata, &sor.Sdata, &sor.Level.Name, &sor.Country.Name, &sor.Subject.Name, &sor.City.Name, &sor.Stadium.Name)
		
		t1, _ := time.Parse("2006-01-02", sor.Fdata)
		t2, _ := time.Parse("2006-01-02", sor.Sdata)
		sor.Fdata = t1.Format("2 January 2006")
		sor.Sdata = t2.Format("2 January 2006")

		send = append(send, sor)
	}
	
	return send
}

func history(id string) []sorevnovanie {
	send := []sorevnovanie {}
	data := time.Now().Format("2006-01-02 15:04")
	res, err := database.Query("SELECT sorevnovania.id, sorevnovania.name, sorevnovania.fdata, sorevnovania.sdata, lvl.name, country.name, subj.name, city.name, stad.name FROM `sorevnovania` JOIN levels AS lvl ON lvl.id = sorevnovania.level_id  JOIN address AS country ON country.id = sorevnovania.country_id JOIN address AS subj ON subj.id = sorevnovania.subject_id JOIN address AS city ON city.id = sorevnovania.city_id JOIN address AS stad ON stad.id = sorevnovania.stadium_id WHERE  sorevnovania.sdata < ? AND sport_id = ?", data, id)
	if err != nil {
		fmt.Println(err)
	}
	for res.Next() {
		sor := sorevnovanie {}
		res.Scan(&sor.Id, &sor.Name, &sor.Fdata, &sor.Sdata, &sor.Level.Name, &sor.Country.Name, &sor.Subject.Name, &sor.City.Name, &sor.Stadium.Name)
		
		t1, _ := time.Parse("2006-01-02", sor.Fdata)
		t2, _ := time.Parse("2006-01-02", sor.Sdata)
		sor.Fdata = t1.Format("2 January 2006")
		sor.Sdata = t2.Format("2 January 2006")

		send = append(send, sor)
	}
	
	return send
}

func UserPage(id string) for_user_page {
	u := for_user_page {}

	// Получаем информацию о пользователе
	s := user {}
	row := database.QueryRow("SELECT users.login, users.command_or_player_id, command.name, command.logo, command.present, sports.name FROM `users` JOIN commands_or_players AS command ON command.id = users.command_or_player_id JOIN sports ON sports.id = command.sports_id WHERE users.id = ?", id)
	row.Scan(&s.Login, &s.Command.Id, &s.Command.Name, &s.Command.Logo, &s.Command.Present, &s.Command.Sport.Name)
	u.Data = s

	// Получаем информацию о достижениях
	dd := []rezults_command {}
	d := rezults_command {}
	res, err := database.Query("SELECT rezults_command.plase, rezults_command.sorevnovanie_id, sorev.name FROM `rezults_command` JOIN sorevnovania AS sorev ON sorev.id = rezults_command.sorevnovanie_id WHERE commands_or_players_id = ?", s.Command.Id)
	if err != nil {
		fmt.Println(err)
	}
	for res.Next() {
		res.Scan(&d.Plase, &d.Sorev.Id, &d.Sorev.Name)
		dd = append(dd, d)
	}
	u.Dost = dd

	return u
}