package main

func main() {
	var user User
	user.Id = 2
	user.Name = "tst2"
	user.Email = "tst@tst.co.jp"
	user.Pass = "1234"

	user.updateUser()

}
