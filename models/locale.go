package models

type Locale struct {
	language       string
	keyboardLayout string
}

//func (locale *Locale) writeLocale() {

//	if locale.language == "PTBR" {
//		data := []byte("LANG=pt_BR.UTF-8")
//		err := ioutil.WriteFile("/etc/locale.conf", data, 0666)
//		check(err)
//	}

//	if locale.keyboardLayout == "br-abnt2" {

//		_ = execute("loadkeys br-abnt2")

//		data := []byte("KEYMAP=br-abnt2" + "\n" +
//			"FONT=lat1-16.psfu.gz")

//		err := ioutil.WriteFile("/etc/vconsole.conf", data, 0666)
//		check(err)
//	}

//	if locale.timezone == "Brasília" {
//		_ = execute("timedatectl set-ntp true")
//		_ = execute("ln -s -f /usr/share/zoneinfo/Brazil/East /etc/localtime")
//		_ = execute("hwclock --systohc --utc")

//	}
//}

//func (locale *Locale) setLocale() *Locale {
//	var language, keyboardLayout, timezone, localeValidation string

//	for {

//		fmt.Println("Qual língua você deseja que o sistema possua?")
//		fmt.Println("[PTBR]")
//		fmt.Scanf("%s", &language)

//		fmt.Println("Existe um padrão de teclado definido para sua linguagem. Você deseja usá-la?")
//		fmt.Println("[br-abnt2]")
//		fmt.Scanf("%s", &keyboardLayout)

//		fmt.Println("Fuso horário")
//		fmt.Println("[Brasília]")
//		fmt.Scanf("%s", &timezone)

//		*locale = Locale{language, keyboardLayout, timezone}

//		locale.printLocale()

//		fmt.Println("Os dados estão corretos?")
//		fmt.Println("[sim não]")
//		fmt.Scanf("%s", &localeValidation)

//		if localeValidation == "sim" {
//			break
//		}

//	}

//	return locale
//}

//func (locale *Locale) printLocale() {

//	fmt.Println("Os dados selecionados foram:" +
//		"Língua : " + locale.language + "\n" +
//		"Padrão de teclado : " + locale.keyboardLayout + "\n" +
//		"Fuso horário : " + locale.timezone)

//}
