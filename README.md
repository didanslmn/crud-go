# CRUD Employee Management dengan Go

Aplikasi sederhana untuk manajemen data karyawan dengan operasi CRUD (Create, Read, Update, Delete) menggunakan:
- Bahasa: Go
- Database: MySQL
- Router: Chi

## Fitur
- Membuat data karyawan baru
- Melihat daftar semua karyawan
- Mengupdate data karyawan
- Menghapus data karyaman
- Validasi form input
- Redirect setelah operasi berhasil

## Prasyarat
- Go 1.18+
- MySQL 5.7+
- Go Modules

## Instalasi

1. Clone repository:
```bash
git clone https://github.com/username/crud-go.git
cd crud-go
```
2. Setup databse
``` bash
CREATE DATABASE crud_go;
USE crud_go;

CREATE TABLE employee (
    id INT PRIMARY KEY AUTO_INCREMENT,
    nama VARCHAR(255) NOT NULL,
    npwp VARCHAR(255) NOT NULL,
    address TEXT NOT NULL
);
```
# Jalankan server di port 8080
go run main.go

# Akses di browser:
http://localhost:8080/employees
