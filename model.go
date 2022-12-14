package main

import (
	// DBを操作するためのライブラリ
	"database/sql"

	// sql3を操作するためのドライバを提供
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Pass  string `json:"pass"`
}

var Db *sql.DB

// DBのコネクションの作成
func init() {
	var err error
	// DBとの接続、第一引数にドライバを第二引数にデータソース名を指定
	Db, err = sql.Open("sqlite3", "./example.sqlite")

	if err != nil {
		panic(err)
	}
}

// ユーザーの新規作成
func (user *User) createUser() (err error) {
	cmd := "INSERT INTO users(name, email, pass) values($1, $2, $3) RETURNING id"
	err = Db.QueryRow(cmd, user.Name, user.Email, user.Pass).Scan(&user.Id)
	return
}

// ユーザーの削除
func (user *User) deleteUser() (err error) {
	cmd := "DELETE FROM users WHERE id = $1"
	_, err = Db.Exec(cmd, user.Id)
	return
}

// ユーザーの編集
func (user *User) updateUser() (err error) {
	cmd := "UPDATE users set name = $1, email = $2, pass = $3 WHERE id = $4"
	_, err = Db.Exec(cmd, user.Name, user.Email, user.Pass, user.Id)
	return
}

// IDからユーザー取得
func GetUserById(id int) (user User, err error) {
	cmd := "SELECT id, name, email FROM users WHERE id = $1"
	err = Db.QueryRow(cmd, id).Scan(&user.Id, &user.Name, &user.Email)
	return
}

// 名前とパスワードからユーザー取得
func (user *User) getUser() {
	cmd := "SELECT id, name, email FROM users WHERE name = $1 AND pass = $2"
	Db.QueryRow(cmd, user.Name, user.Pass).Scan(&user.Id, &user.Name, &user.Email)
	return
}
