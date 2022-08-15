package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	// アドレスポートの指定
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/users/create", requestCreate)

	// サーバーの起動
	server.ListenAndServe()

}

// ユーザーの作成
func requestCreate(w http.ResponseWriter, r *http.Request) {
	// リクエストのContent-Lengthを取得
	contentLength := r.ContentLength
	// リクエストのbodyを格納するために取得した長さのbyte型スライスcontentBodyを作成
	contentBody := make([]byte, contentLength)
	// リクエストを格納
	r.Body.Read(contentBody)

	var user User
	// jsonの詰め替え
	err := json.Unmarshal(contentBody, &user)
	if err != nil {
		panic(err)
	}

	// 既に登録されていないか確認
	user.getUser()
	if user.Id != 0 {
		panic("既に登録されているユーザー")
	}

	err = user.createUser()
	if err != nil {
		panic(err)
	}

	// postを再度jsonに変換してレスポンスとして返す
	output, err := json.MarshalIndent(&user, "", "\t")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(200)
	w.Write(output)
	return
}
