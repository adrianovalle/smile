// smile

package main

import (
	"fmt"
	"io/ioutil"
	//	"os"
	"os/exec"
	"regexp"
	"strings"
)

const version = "0.4.0"

var verbose bool

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func execute(cmd string) []byte {

	cmdLine := strings.Fields(cmd)
	command := cmdLine[0]
	parameters := cmdLine[1:len(cmdLine)]

	out, err := exec.Command(command, parameters...).Output()

	check(err)
	if verbose == true {
		fmt.Printf("%s", out)
	}
	return out
}

type ConnectionProfile struct {
	wifiInterface    string
	connectionType   string
	wifiSecurityType string
	essid            string
	ipMode           string
	wifiPassword     string
	hidden           string
}

func (connProfile *ConnectionProfile) writeWifiConfigToFile(destinationFolder string) {

	const description = "Description=Configuration file created by smile " + version
	const descWifiInterface = "Interface="
	const descConnectionType = "Connection="
	const descWifiSecurityType = "Security="
	const descEssid = "ESSID="
	const descIpMode = "IP="
	const descWifiPass = "Key="
	const descHidden = "Hidden="

	data := []byte(description + "\n" +

		descWifiInterface + connProfile.wifiInterface + "\n" +

		descConnectionType + connProfile.connectionType + "\n" +

		descWifiSecurityType + connProfile.wifiSecurityType + "\n" +

		descEssid + connProfile.essid + "\n" +

		descIpMode + connProfile.ipMode + "\n" +

		descWifiPass + connProfile.wifiPassword + "\n" +

		descHidden + connProfile.hidden)

	err := ioutil.WriteFile(destinationFolder+"/"+connProfile.essid, data, 0777)
	check(err)

}
func detectNetwork() []string {

	cmdOut := execute("ip link")

	r, _ := regexp.Compile("[a-z]{3}[0-9][a-z][0-9]")

	regex := r.FindAllString(string(cmdOut), -1)
	return regex
}
func detectPartitionTable() []string {

	cmdOut := execute("lsblk")
	r, _ := regexp.Compile("[s]{1}[d]{1}[a-z]{1}[ ]{1}")
	regex := r.FindAllString(string(cmdOut), -1)
	return regex
}

func (connProfile *ConnectionProfile) printConnectionProfile() {
	fmt.Printf("%s", execute("clear"))
	fmt.Println("\t\t\t\t\tOs dados informados foram: \n\n" +

		"\t\t\t\t\tInterface de rede: " + connProfile.wifiInterface + "\n\n" +

		"\t\t\t\t\tTipo de Conexao: " + connProfile.connectionType + "\n\n" +

		"\t\t\t\t\tSeguranca da rede: " + connProfile.wifiSecurityType + "\n\n" +

		"\t\t\t\t\tNome da Rede Wi-fi: " + connProfile.essid + "\n\n" +

		"\t\t\t\t\tModo de aquisicao IP: " + connProfile.ipMode + "\n\n" +

		"\t\t\t\t\tSenha da Rede: " + connProfile.wifiPassword + "\n\n" +

		"\t\t\t\t\tRede Oculta: " + connProfile.hidden + "\n")

}

func (connProfile *ConnectionProfile) setConnectionProfile() *ConnectionProfile {
	var wifiInterface, connectionType, wifiSecurityType, essid, ipMode, wifiPassword, hidden, wifiProfileValidation string

	for {
		fmt.Printf("%s", execute("clear"))

		fmt.Printf("Informe sua interface de rede \n")
		fmt.Println(detectNetwork())
		fmt.Scanf("%s", &wifiInterface)

		fmt.Printf("Informe o tipo de conexao \n")
		fmt.Println("[wireless ethernet]")
		fmt.Scanf("%s", &connectionType)

		fmt.Printf("Informe a seguranca wi-fi \n")
		fmt.Println("[wpa wpa2]")
		fmt.Scanf("%s", &wifiSecurityType)

		fmt.Printf("Informe a identificação da rede \n")
		fmt.Scanf("%s", &essid)

		fmt.Printf("Informe o modo de Ip \n")
		fmt.Println("[dhcp]")
		fmt.Scanf("%s", &ipMode)

		fmt.Printf("Informe a senha da rede \n")
		fmt.Scanf("%s", &wifiPassword)

		fmt.Printf("A rede está oculta? \n")
		fmt.Println("[yes no]")
		fmt.Scanf("%s", &hidden)

		*connProfile = ConnectionProfile{wifiInterface, connectionType, wifiSecurityType, essid, ipMode, wifiPassword, hidden}

		connProfile.printConnectionProfile()

		fmt.Println("Os dados estao corretos?")
		fmt.Println("[yes no]")
		fmt.Scanf("%s", &wifiProfileValidation)

		if wifiProfileValidation == "yes" {
			break
		}

	}

	return connProfile
}

type Locale struct {
	language       string
	keyboardLayout string
	timezone       string
}

func (locale *Locale) writeLocale() {

	if locale.language == "Português-Brasileiro" {
		data := []byte("LANG=pt_BR.ISO-8859-1")
		err := ioutil.WriteFile("/etc/locale.conf", data, 0777)
		check(err)
	}

	if locale.keyboardLayout == "[br-abnt2]" {

		_ = execute("loadkeys br-abnt2")

		data := []byte("KEYMAP=br.abnt2" + "\n" +
			"FONT=lat1-16.psfu.gz")

		err := ioutil.WriteFile("/etc/vconsole.conf", data, 0777)
		check(err)
	}

	if locale.timezone == "Brasilia" {
		_ = execute("timedatectl set-ntp true")
		_ = execute("ln -s -f /usr/share/zoneinfo/America/Sao_Paulo /etc/localtime")

	}
}

func (locale *Locale) setLocale() *Locale {
	var language, keyboardLayout, timezone, localeValidation string

	for {

		fmt.Println("Qual língua você deseja que o sistema posa?")
		fmt.Println("[Português-Brasileiro]")
		fmt.Scanf("%s", &language)

		fmt.Println("Existe um padrão de teclado definido para sua linguagem. Você desa usa-la?")
		fmt.Println("[br-abnt2]")
		fmt.Scanf("%s", &keyboardLayout)

		fmt.Println("Fuso horario")
		fmt.Println("[Brasilia]")
		fmt.Scanf("%s", &timezone)

		*locale = Locale{language, keyboardLayout, timezone}

		locale.printLocale()

		fmt.Println("Os dados estao corretos?")
		fmt.Println("[yes no]")
		fmt.Scanf("%s", &localeValidation)

		if localeValidation == "yes" {
			break
		}

	}

	return locale
}

func (locale *Locale) printLocale() {

	fmt.Println("Os dados selecionados foram:" +
		"Lingua : " + locale.language + "\n" +
		"Padrao de teclado : " + locale.keyboardLayout + "\n" +
		"Fuso horario : " + locale.timezone)

}

type Partition struct {
	device     string
	filesystem string
	efiSupport string
}

func (partition *Partition) setPartition() {
	var filesystem, efiSupport, partitionValidation string

	for {
		fmt.Println("Informe em qual dispositivo voce deseja criar o particionamento")
		fmt.Println(detectPartitionTable())
		fmt.Scanf("%s", &filesystem)

		fmt.Println("Seu computador tem suporte a EFI?")
		fmt.Prinln("[yes no]")
		fmt.Scanf("%s", &efiSupport)

		*partition = Partition{device, filesystem, efiSupport}

		


		fmt.Println ("Os dados est�o corretos?")
		fmt.Scanf("%s",partitionValidation)

		if partitionValidation == "yes" {
			break
		}

	}

}

func (partition *Partition) printPartition() {

		fmt.Println("Parti��o selecionada: " + partition.device + "/n" +
			    "Sistema de arquivos:: " + partition.filesystem + "/n"" +
			    UEFI : " + partition.efiSupport)


}







//if efiSupport == "no" {

//}

//	fmt.Println("Você deseja utilizar todo o espaço da partição para o sistema?")
//	fmt.Println("[yes]")
//	fmt.Scanf("%s",&)
//createPartitionTable()

//parted

//	fmt.Println("O sistema será instalado em um pendrive ou ssd?")
//	fmt.Println("[yes]")
//	fmt.Scanf("%s",&)

//	fmt.Println("Você deseja instalar então o sistema de arquivos f2fs, que permite um melhor aproveitamento para esse tipo de disposivo?")
//	fmt.Println("[yes]")
//	fmt.Scanf("%s",&)
//mkfs.f2fs

//}

func main() {
	var connProfile ConnectionProfile
	//	var efiSupport string
	var locale Locale

	verbose = false

	//teclado

	locale = *locale.setLocale()

	locale.writeLocale()

	//conexão de rede
	connProfile = *connProfile.setConnectionProfile()

	connProfile.writeWifiConfigToFile("/etc/netctl")

	_ = execute("netctl start " + connProfile.essid) //substitui o wifi-menu

	//particionamento

	//mkdir -p /mnt/boot
	//mount ????

	//pacstrap -i /mnt base base-devel
	//genfstab -U /mnt > /mnt/etc/fstab
	//arch-chroot /mnt /bin/bash

	//tzselect
	//ln -sf /usr/share/zoneinfo/Zone???/Subzone??? /etc/localtime

	//hwclock --systohc --utc
	//mkinitcpio -p linux

}
