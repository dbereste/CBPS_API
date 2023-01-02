package invoice

import (
	"database/sql"
	"fmt"
	"main/config"
	userr "main/models/user"
)

type InvoicesView struct {
	Id           int
	Name         string
	Surname      string
	Date         string
	DueDate      string
	Items        []InvoicesViewDetails
	Total        int
	Email        string
	BusinessName string
}

type InvoicesViewDetails struct {
	InvoiceId int
	Name      string
	Price     int
}

type DataDetails struct {
	Items []InvoicesViewDetails
}

type Invoices struct {
	Items []InvoicesView
}

func GetAll(user string) Invoices {
	data := Invoices{}
	db, err := sql.Open("mysql", config.DBuser+":"+config.DBpass+"@tcp(127.0.0.1:3306)/"+config.DBname)

	if err != nil {
		fmt.Println(err)
	}

	userdb := userr.GetBylogin(user)

	sqll := "SELECT t1.id as `id`, t2.name as `name`, t2.surname as `surname`, t1.`created_at` as `date`, t1.due_date as `due_date`, t2.email as `email`, t2.business_name as `business_name` FROM cbps.invoice_items t1 INNER JOIN cbps.users t2 on t1.user_id = t2.id"

	if userdb.Level == 0 {
		sqll = sqll + " WHERE t2.login = '" + user + "'"
	}

	row, err := db.Query(sqll)

	defer row.Close()

	for row.Next() {
		invoice := InvoicesView{}
		err = row.Scan(&invoice.Id, &invoice.Name, &invoice.Surname, &invoice.Date, &invoice.DueDate, &invoice.Email, &invoice.BusinessName)

		sql1 := "SELECT invoice_id, name, price FROM cbps.invoice_items_details WHERE invoice_id = " + fmt.Sprint(invoice.Id)
		row1, err1 := db.Query(sql1)
		if err1 != nil {
			fmt.Println(err)
		}
		defer row1.Close()

		datadetails := DataDetails{}
		var Sum int = 0
		for row1.Next() {

			invoice_details := InvoicesViewDetails{}
			err1 = row1.Scan(&invoice_details.InvoiceId, &invoice_details.Name, &invoice_details.Price)
			Sum += invoice_details.Price
			datadetails.Items = append(datadetails.Items, invoice_details)

		}
		invoice.Total = Sum
		invoice.Items = datadetails.Items
		data.Items = append(data.Items, invoice)
	}

	defer db.Close()
	return data
}

func GetPaid(user string) Invoices {
	data := Invoices{}
	db, err := sql.Open("mysql", config.DBuser+":"+config.DBpass+"@tcp(127.0.0.1:3306)/"+config.DBname)

	if err != nil {
		fmt.Println(err)
	}
	userdb := userr.GetBylogin(user)

	sqll := "SELECT t1.id as `id`, t2.name as `name`, t2.surname as `surname`, t1.`created_at` as `date`, t1.due_date as `due_date`, t2.email as `email`, t2.business_name as `business_name`  FROM cbps.invoice_items t1 INNER JOIN cbps.users t2 on t1.user_id = t2.id WHERE t1.status = 1"

	if userdb.Level == 0 {
		sqll = sqll + " AND t2.login = '" + user + "'"
	}

	row, err := db.Query(sqll)

	defer row.Close()

	for row.Next() {
		invoice := InvoicesView{}
		err = row.Scan(&invoice.Id, &invoice.Name, &invoice.Surname, &invoice.Date, &invoice.DueDate, &invoice.Email, &invoice.BusinessName)

		sql1 := "SELECT invoice_id, name, price FROM cbps.invoice_items_details WHERE invoice_id = " + fmt.Sprint(invoice.Id)
		row1, err1 := db.Query(sql1)

		if err1 != nil {
			fmt.Println(err)
		}
		defer row1.Close()

		datadetails := DataDetails{}
		var Sum int = 0
		for row1.Next() {

			invoice_details := InvoicesViewDetails{}
			err1 = row1.Scan(&invoice_details.InvoiceId, &invoice_details.Name, &invoice_details.Price)
			Sum += invoice_details.Price
			datadetails.Items = append(datadetails.Items, invoice_details)

		}
		invoice.Total = Sum
		invoice.Items = datadetails.Items
		data.Items = append(data.Items, invoice)
	}

	defer db.Close()

	return data
}

func GetUnpaid(user string) Invoices {
	data := Invoices{}
	db, err := sql.Open("mysql", config.DBuser+":"+config.DBpass+"@tcp(127.0.0.1:3306)/"+config.DBname)

	if err != nil {
		fmt.Println(err)
	}
	userdb := userr.GetBylogin(user)

	sqll := "SELECT t1.id as `id`, t2.name as `name`, t2.surname as `surname`, t1.`created_at` as `date`, t1.due_date as `due_date`, t2.email as `email`, t2.business_name as `business_name`  FROM cbps.invoice_items t1 INNER JOIN cbps.users t2 on t1.user_id = t2.id WHERE t1.status = 0"

	if userdb.Level == 0 {
		sqll = sqll + " AND t2.login = '" + user + "'"
	}

	row, err := db.Query(sqll)

	defer row.Close()

	for row.Next() {
		invoice := InvoicesView{}
		err = row.Scan(&invoice.Id, &invoice.Name, &invoice.Surname, &invoice.Date, &invoice.DueDate, &invoice.Email, &invoice.BusinessName)

		sql1 := "SELECT invoice_id, name, price FROM cbps.invoice_items_details WHERE invoice_id = " + fmt.Sprint(invoice.Id)
		row1, err1 := db.Query(sql1)

		if err1 != nil {
			fmt.Println(err)
		}
		defer row1.Close()

		datadetails := DataDetails{}
		var Sum int = 0
		for row1.Next() {

			invoice_details := InvoicesViewDetails{}
			err1 = row1.Scan(&invoice_details.InvoiceId, &invoice_details.Name, &invoice_details.Price)
			Sum += invoice_details.Price
			datadetails.Items = append(datadetails.Items, invoice_details)

		}
		invoice.Total = Sum
		invoice.Items = datadetails.Items
		data.Items = append(data.Items, invoice)
	}

	defer db.Close()

	return data
}

func GetByID(id, user string) InvoicesView {
	data := InvoicesView{}
	db, err := sql.Open("mysql", config.DBuser+":"+config.DBpass+"@tcp(127.0.0.1:3306)/"+config.DBname)

	if err != nil {
		fmt.Println(err)
	}

	userdb := userr.GetBylogin(user)

	sqll := "SELECT t1.id as `id`, t2.name as `name`, t2.surname as `surname`, t1.`created_at` as `date`, t1.due_date as `due_date`, t2.email as `email`, t2.business_name as `business_name`  FROM cbps.invoice_items t1 INNER JOIN cbps.users t2 on t1.user_id = t2.id WHERE t1.id = '" + id + "'"
	row, err := db.Query(sqll)

	if userdb.Level == 0 {
		sqll = sqll + "t2.login = '" + user + "' and"
	}

	defer row.Close()

	for row.Next() {
		invoice := InvoicesView{}
		err = row.Scan(&invoice.Id, &invoice.Name, &invoice.Surname, &invoice.Date, &invoice.DueDate, &invoice.Email, &invoice.BusinessName)

		sql1 := "SELECT invoice_id, name, price FROM cbps.invoice_items_details WHERE invoice_id = " + fmt.Sprint(invoice.Id)
		row1, err1 := db.Query(sql1)

		if err1 != nil {
			fmt.Println(err)
		}
		defer row1.Close()

		datadetails := DataDetails{}
		var Sum int = 0
		for row1.Next() {

			invoice_details := InvoicesViewDetails{}
			err1 = row1.Scan(&invoice_details.InvoiceId, &invoice_details.Name, &invoice_details.Price)
			Sum += invoice_details.Price
			datadetails.Items = append(datadetails.Items, invoice_details)

		}
		invoice.Total = Sum
		invoice.Items = datadetails.Items

		data = invoice
	}

	defer db.Close()

	return data
}

func New(usernew InvoicesView) int {

	return 0
}

func Update(useredit InvoicesView) error {

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
