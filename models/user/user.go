package user

import (
	"database/sql"
	"fmt"
	"main/config"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id            int
	Email         string
	Password      string
	Name          string
	Surname       string
	Login         string
	Lastlogindate string
	Business_name string
	Address       string
	Phone_number  int
	Level         int
}

type Users struct {
	Users []User
}

type UserLogin struct {
	Login    string
	Password string
}

func GetLoginInfo(user string) (string, string) {
	var data UserLogin

	db, err := sql.Open("mysql", config.DBuser+":"+config.DBpass+"@tcp(127.0.0.1:3306)/"+config.DBname)

	if err != nil {
		fmt.Println(err)
	}

	sqll := "SELECT login, password FROM cbps.users WHERE login = '" + user + "'"

	errr := db.QueryRow(sqll).Scan(&data.Login, &data.Password)

	if errr != nil {
		fmt.Println(errr)
	}

	return data.Login, data.Password
}

func GetAll() Users {
	var data Users

	db, err := sql.Open("mysql", config.DBuser+":"+config.DBpass+"@tcp(127.0.0.1:3306)/"+config.DBname)

	if err != nil {
		fmt.Println(err)
	}

	sqll := "SELECT id, email, password, name , surname, login, lastlogindate, business_name, address, phone_number FROM cbps.users"

	rows, err := db.Query(sqll)

	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()

	for rows.Next() {
		var tmpUser User
		err = rows.Scan(&tmpUser.Id, &tmpUser.Email, &tmpUser.Password, &tmpUser.Name, &tmpUser.Surname, &tmpUser.Login, &tmpUser.Lastlogindate, &tmpUser.Business_name, &tmpUser.Address, &tmpUser.Phone_number)
		data.Users = append(data.Users, tmpUser)
	}

	return data
}

func GetByID(id string) User {
	var data User

	db, err := sql.Open("mysql", config.DBuser+":"+config.DBpass+"@tcp(127.0.0.1:3306)/"+config.DBname)

	if err != nil {
		fmt.Println(err)
	}

	sqll := "SELECT id, email, password, name , surname, login, lastlogindate, business_name, address, phone_number, level FROM cbps.users WHERE id = " + id

	rows, err := db.Query(sqll)

	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&data.Id, &data.Email, &data.Password, &data.Name, &data.Surname, &data.Login, &data.Lastlogindate, &data.Business_name, &data.Address, &data.Phone_number, &data.Level)
	}

	return data
}

func GetBylogin(login string) User {
	var data User

	db, err := sql.Open("mysql", config.DBuser+":"+config.DBpass+"@tcp(127.0.0.1:3306)/"+config.DBname)

	if err != nil {
		fmt.Println(err)
	}

	sqll := "SELECT id, email, password, name , surname, login, lastlogindate, business_name, address, phone_number, level FROM cbps.users WHERE login = '" + login + "'"

	rows, err := db.Query(sqll)

	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&data.Id, &data.Email, &data.Password, &data.Name, &data.Surname, &data.Login, &data.Lastlogindate, &data.Business_name, &data.Address, &data.Phone_number, &data.Level)
	}

	return data
}

func New(usernew User) int {
	db, err := sql.Open("mysql", config.DBuser+":"+config.DBpass+"@tcp(127.0.0.1:3306)/"+config.DBname)

	if err != nil {
		fmt.Println(err)
		return 0
	}
	var lastInsertId int64
	sqll := "INSERT INTO cbps.users (email, password, name, surname, login, business_name, address, phone_number) VALUES ('" + usernew.Email + "', '" + usernew.Password + "', '" + usernew.Name + "', '" + usernew.Surname + "', '" + usernew.Login + "', '" + usernew.Business_name + "', '" + usernew.Address + "'," + strconv.Itoa(usernew.Phone_number) + ")"

	res, err := db.Exec(sqll)

	// TODO: Write proper Error handler
	if err != nil {
		fmt.Println(err)
		return 0
	}

	lastInsertId, err = res.LastInsertId()

	if err != nil {
		fmt.Println(err)
		return 0
	}

	return int(lastInsertId)
}

func Update(useredit User) error {
	db, err := sql.Open("mysql", config.DBuser+":"+config.DBpass+"@tcp(127.0.0.1:3306)/"+config.DBname)

	if err != nil {
		fmt.Println(err)
		return err
	}

	sqll := "UPDATE cbps.users SET email = '" + useredit.Email + "', password = '" + useredit.Password + "', name = '" + useredit.Password + "', surname = '" + useredit.Surname + "', login = '" + useredit.Login + "', lastlogindate = '" + useredit.Lastlogindate + "', business_name = '" + useredit.Business_name + "', address = '" + useredit.Address + "', phone_number = " + fmt.Sprint(useredit.Phone_number) + " WHERE id = " + fmt.Sprint(useredit.Id)

	res, err := db.Exec(sqll)

	fmt.Println(res)

	// TODO: Write proper Error handler
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func Delete(id string) int {

	db, err := sql.Open("mysql", config.DBuser+":"+config.DBpass+"@tcp(127.0.0.1:3306)/"+config.DBname)

	if err != nil {
		fmt.Println(err)
	}

	sqll := "DELETE FROM cbps.users WHERE id = " + id

	rows, err := db.Query(sqll)

	if err != nil {
		fmt.Println(err)
	}

	defer rows.Close()

	return 0
}
