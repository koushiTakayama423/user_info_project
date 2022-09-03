package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"path"
	"strconv"
)

func main() {
	// アドレスポートの指定
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/users", handleRequest)

	// サーバーの起動
	server.ListenAndServe()

}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	var err error

	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

	switch r.Method {
	case "GET":
		err = errors.New("Getメソッドは未対応")
	case "POST":
		err = requestCreate(w, r)
	case "PUT":
		err = requestEdit(w, r)
	case "DELETE":
		err = requestDelete(w, r)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// ユーザーの作成
func requestCreate(w http.ResponseWriter, r *http.Request) (err error) {
	var user User
	err = user.readJson(r)
	if err != nil {
		return
	}

	err = user.userChecker()
	if err != nil {
		return
	}

	// 既に登録されていないか確認
	user.getUser()
	if user.Id != 0 {
		err = errors.New("ユーザーが重複")
		return
	}

	err = user.createUser()
	if err != nil {
		return
	}

	// postを再度jsonに変換してレスポンスとして返す
	output, err := json.MarshalIndent(&user, "", "\t")
	if err != nil {
		return
	}

	w.WriteHeader(200)
	w.Write(output)
	return
}

// ユーザーの編集
func requestEdit(w http.ResponseWriter, r *http.Request) (err error) {
	var user User
	err = user.readJson(r)
	if err != nil {
		return
	}

	err = user.userChecker()
	if err != nil {
		return
	}

	// 既に登録されていないか確認
	user.getUser()
	if user.Id != 0 {
		err = errors.New("ユーザーが重複")
		return
	}

	err = user.updateUser()
	if err != nil {
		return
	}

	// postを再度jsonに変換してレスポンスとして返す
	output, err := json.MarshalIndent(&user, "", "\t")
	if err != nil {
		return
	}

	w.WriteHeader(200)
	w.Write(output)
	return
}

// ユーザーの削除
func requestDelete(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}

	user, err := GetUserById(id)
	if err != nil {
		return
	}

	user.deleteUser()

	// postを再度jsonに変換してレスポンスとして返す
	output, err := json.MarshalIndent(&user, "", "\t")
	if err != nil {
		return
	}

	w.WriteHeader(200)
	w.Write(output)
	return
}

// ユーザー情報のチェック
func (user *User) userChecker() (err error) {
	if user.Name == "" {
		err = errors.New("ユーザー名が空")
		return
	} else if user.Email == "" {
		err = errors.New("ユーザーメールアドレスが空")
		return
	} else if user.Pass == "" {
		err = errors.New("ユーザーパスが空")
		return
	}
	return
}

// リクエストjsonをuserに格納する
func (user *User) readJson(r *http.Request) (err error) {
	// リクエストのContent-Lengthを取得
	contentLength := r.ContentLength
	// リクエストのbodyを格納するために取得した長さのbyte型スライスcontentBodyを作成
	contentBody := make([]byte, contentLength)
	// リクエストを格納
	r.Body.Read(contentBody)
	fmt.Println(contentBody)

	// jsonの詰め替え
	err = json.Unmarshal(contentBody, &user)
	return
}
