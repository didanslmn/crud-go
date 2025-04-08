package model

import (
	"database/sql"
	"fmt"
)

// Struct Employee merepresentasikan tabel database employee
type Employee struct {
	ID      int
	Nama    string
	NPWP    string
	Address string
}

func (e *Employee) Save(db *sql.DB) error {
	// Validasi field wajib diisi
	if e.Nama == "" || e.NPWP == "" || e.Address == "" {
		return fmt.Errorf("semua field harus diisi")
	}

	// Eksekusi query INSERT dan gunakan parameterized query untuk prevent SQL injection
	res, err := db.Exec("INSERT INTO employee(nama, npwp, address) VALUE (?,?,?)", e.Nama, e.NPWP, e.Address)
	if err != nil {
		return fmt.Errorf("gagal insert data: %w", err) // Wrap error dengan pesan lebih jelas
	}
	id, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("gagal mendapatkan ID terakhir: %w", err)
	}
	e.ID = int(id)
	return nil
}

func GetAll(db *sql.DB) ([]Employee, error) {
	// Eksekusi query SELECT semua data
	rows, err := db.Query("SELECT id, nama, npwp, address FROM employee")
	if err != nil {
		return nil, err
	}
	defer rows.Close() // rows ditutup setelah selesai

	var employees []Employee
	// Iterasi melalui setiap row hasil query
	for rows.Next() {
		var e Employee
		// Map kolom database ke struct employee
		err = rows.Scan(&e.ID, &e.Nama, &e.NPWP, &e.Address)
		if err != nil {
			return nil, err
		}
		employees = append(employees, e)
	}
	return employees, nil
}

func GetByID(db *sql.DB, id int) (*Employee, error) {
	var e Employee
	// Gunakan QueryRow untuk single row result
	err := db.QueryRow("SELECT id, nama, npwp, address FROM employee WHERE id=?", id).Scan(&e.ID, &e.Nama, &e.NPWP, &e.Address)
	if err != nil {
		return nil, err
	}
	return &e, nil
}

func (e *Employee) Update(db *sql.DB) error {
	// Eksekusi query update
	_, err := db.Exec("UPDATE employee SET nama=?, npwp=?, address=? WHERE id=?", e.Nama, e.NPWP, e.Address, e.ID)
	return err
}

func Delete(db *sql.DB, id int) error {
	// Eksekusi uery delete
	_, err := db.Exec("DELETE FROM employee WHERE id = ?", id)
	return err
}
