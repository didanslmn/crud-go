package model

import (
	"database/sql"
	"fmt"
)

type Employee struct {
	ID      int
	Nama    string
	NPWP    string
	Address string
}

func (e *Employee) Save(db *sql.DB) error {
	if e.Nama == "" || e.NPWP == "" || e.Address == "" {
		return fmt.Errorf("semua field harus diisi")
	}

	res, err := db.Exec("INSERT INTO employee(nama, npwp, address) VALUE (?,?,?)", e.Nama, e.NPWP, e.Address)
	if err != nil {
		return err
	}
	id, _ := res.LastInsertId()
	e.ID = int(id)
	return nil
}

func GetAll(db *sql.DB) ([]Employee, error) {
	rows, err := db.Query("SELECT id, nama, npwp, address FROM employee")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []Employee
	for rows.Next() {
		var e Employee
		err = rows.Scan(&e.ID, &e.Nama, &e.NPWP, &e.Address)
		if err != nil {
			return nil, err
		}
		employees = append(employees, e)
	}
	return employees, nil
}
