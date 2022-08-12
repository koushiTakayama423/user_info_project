package main

import (
	// DBを操作するためのライブラリ
	"database/sql"
	"fmt"

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
	if user.Name == "" {
		panic("ユーザー名が空")
	} else if user.Email == "" {
		panic("ユーザーメールアドレスが空")
	} else if user.Pass == "" {
		panic("ユーザーパスが空")
	}

	cmd := "INSERT INTO users(name, email, pass) values($1, $2, $3) RETURNING id"
	err = Db.QueryRow(cmd, user.Name, user.Email, user.Pass).Scan(&user.Id)
	fmt.Println(err)
	return
}