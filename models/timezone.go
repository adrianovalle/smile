package models

type Timezone struct {
	location string
}

//func (timezone *Timezone) setTimezone() {

//	if locale.timezone == "Brasília" {
//		_ = execute("timedatectl set-ntp true")
//		_ = execute("ln -s -f /usr/share/zoneinfo/Brazil/East /etc/localtime")
//		_ = execute("hwclock --systohc --utc")

//	}
//}

//func (locale *Locale) setLocale() *Locale {
//	var language, keyboardLayout, timezone, localeValidation string

//	for {

//		fmt.Println("Fuso horário")
//		fmt.Println("[Brasília]")
//		fmt.Scanf("%s", &timezone)

//		*locale = Locale{language, keyboardLayout, timezone}

//	}

//	return locale
//}
