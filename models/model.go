package models

import (
	// DBを操作するためのライブラリ
	"database/sql"
	"fmt"

	// sql3を操作するためのドライバを提供
	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

type Tst struct {
	Id  int
	Txt string
}

// DBのコネクションの作成
func DbConnection() {
	var err error
	// DBとの接続、第一引数にドライバを第二引数にデータソース名を指定
	Db, err = sql.Open("sqlite3", "./example.sqlite")

	if err != nil {
		panic(err)
	}
}

func Dbtst() {
	cmd := "select * from tst"
	row := Db.QueryRow(cmd)
	var t Tst
	row.Scan(&t.Id, &t.Txt)
	fmt.Println(&t)
}
