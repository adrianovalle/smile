package models

type User struct {
	username string
	password string
}

//func addUser() {
//	var username string
//	fmt.Println("Digite o nome do usuario")
//	fmt.Scanf("%s", &username)

//	_ = executeInArchChroot("'useradd' -m -s /bin/bash -G wheel,users,audio,video,input,games " + username)
//	enableUserToUseSudo()
//}

//func setPassword(user string) {
//	fmt.Println("Digite uma senha para o usuário " + user)
//	_ = execute("passwd " + user)

//}
