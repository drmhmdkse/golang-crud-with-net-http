package databaseOP

import (
	"database/sql"
	"fmt"
	"log"
)

func ConnectDb() (*sql.DB, error) {
	vt, err := sql.Open("mysql", "root:1965@(127.0.0.1:3306)/users?parseTime=true")

	if err != nil {
		return nil, err
	}

	return vt, nil
}

func Get(id int) *sql.Row {
	vt, err := ConnectDb()

	if err != nil {
		log.Println("connect hatası: ", err)
	}
	sorgu := "SELECT username, firstname, password FROM user where id=?"
	row := vt.QueryRow(sorgu, id)
	defer vt.Close()
	return row
}

func Create(username string, firstname string, password string) sql.Result {
	vt, err := ConnectDb()
	if err != nil {
		log.Println("connect hatası:", err)
	}
	sorgu := "INSERT INTO user (username,firstname,password) VALUES (?,?,?)"

	sonuc, hata := vt.Exec(sorgu, username, firstname, password)
	if hata != nil {
		log.Println("create edilemedi : ", hata)
	}
	defer vt.Close()
	return sonuc
}

func Delete(id int) sql.Result {
	vt, err := ConnectDb()
	if err != nil {
		log.Println("connect hatası:", err)
		return nil
	}
	sonuc, _ := vt.Exec("DELETE FROM user WHERE id=?", id)
	defer vt.Close()
	return sonuc
}


func Update(id int, username string, firstname string, password string) sql.Result {
	vt, err := ConnectDb()
	if err != nil {
		log.Println("burada hata var", err)
	}
	sorguCumlecigi := "UPDATE user SET username=?, firstname=?,password=? WHERE id=?"

	sonuc, hata := vt.Exec(sorguCumlecigi, username, firstname, password, id)

	if hata != nil {
		log.Fatalln("okunamadı satır 66: ", hata)
	}
	defer vt.Close()
	return sonuc
}

func Exist(username string) bool { // connect işlemini zaten yapıyoruz düzenleme yap
	var count int
	vt, _ := ConnectDb()
	sorgu := "SELECT COUNT(*) FROM user WHERE username = ?"

	err := vt.QueryRow(sorgu, username).Scan(&count)
	fmt.Println(count)
	if err != nil {
		log.Println("satır 79: ", err)
	}
	defer vt.Close()
	return count > 0
}

//struct with user data but id is auto increment
