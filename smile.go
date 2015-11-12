// smile


package main

import (
	"fmt"
	"os/exec"
	"strings"
	"io/ioutil"
	"os"
	"regexp"
//	"strconv"
)

const version ="0.3.0"
var verbose bool

func check(e error) {
	if e != nil {
		panic(e)
	}	
}

func execute(cmd string) []byte{

	
	cmdLine := strings.Fields(cmd)
	command := cmdLine[0]
	parameters := cmdLine[1:len(cmdLine)]

	out, err := exec.Command(command, parameters...).Output()

	check(err)
	if verbose==true{
		fmt.Printf("%s", out)
	}
	return out
}

type ConnectionProfile struct{

	wifiInterface    string
	connectionType   string	
	wifiSecurityType string
	essid		 string
	ipMode		 string
	wifiPassword	 string
	hidden		 string 


}

func (connProfile *ConnectionProfile) writeWifiConfigToFile(destinationFolder string, nameProfile string) {

	const description = "Description=Configuration file created by smile " + version
	const descWifiInterface = "Interface="
	const descConnectionType = "Connection="
	const descWifiSecurityType = "Security="
	const descEssid = "ESSID="
	const descIpMode = "IP="
	const descWifiPass = "Key="
	const descHidden = "Hidden="


	data:=[]byte(description +"\n"+

		     descWifiInterface + connProfile.wifiInterface +"\n"+

		     descConnectionType + connProfile.connectionType +"\n"+

		     descWifiSecurityType + connProfile.wifiSecurityType +"\n"+

		     descEssid + connProfile.essid +"\n"+

		     descIpMode + connProfile.ipMode +"\n"+

		     descWifiPass + connProfile.wifiPassword +"\n"+

		     descHidden + connProfile.hidden)


	err:=ioutil.WriteFile(destinationFolder + "/" + nameProfile,data,0777)
	check(err)


}
func copyFile(originalFileWithPath string, destinyFileWithPath string){

	err := os.Rename(originalFileWithPath, destinyFileWithPath)

	check(err)

}
func detectNetwork() []string{


	cmdOut:=execute("ip link")
	
	r,_:=regexp.Compile("[a-z]{3}[0-9][a-z][0-9]")

	regex:=r.FindAllString(string(cmdOut), -1)
	return regex
}
func detectPartitionTable() []string{

	cmdOut := execute("lsblk")
	r,_ := regexp.Compile("[s]{1}[d]{1}[a-z]{1}[ ]{1}")
	regex := r.FindAllString(string(cmdOut),-1)
	return regex
}

func main() {
var wifiInterface, connectionType, wifiSecurityType, essid,ipMode, wifiPassword, hidden string



	verbose=false
	fmt.Printf("%s",execute("clear"))

	fmt.Printf("Bom dia! Informe sua interface de rede \n")
	fmt.Println(detectNetwork())
	fmt.Scanf("%s",&wifiInterface)

	fmt.Printf("Informe o tipo de conexão \n")
	fmt.Println("[wireless ethernet]") 
	fmt.Scanf("%s" ,&connectionType)
	

	fmt.Printf("Informe a segurança wi-fi \n")
	fmt.Println("[wpa wpa2]")
	fmt.Scanf("%s", &wifiSecurityType)


	fmt.Printf("Informe a identificação da rede \n")
	fmt.Scanf("%s" , &essid)


	fmt.Printf("Informe o modo de Ip \n")
	fmt.Println("[ dhcp ]")
	fmt.Scanf("%s", &ipMode)


	fmt.Printf("Informe a senha da rede \n")
	fmt.Scanf("%s", &wifiPassword)


	fmt.Printf("A rede está oculta? \n")
	fmt.Println("[yes no]")
	fmt.Scanf("%s", &hidden)


	connProfile := ConnectionProfile{wifiInterface, connectionType, wifiSecurityType, essid, ipMode, wifiPassword, hidden}
	fmt.Println(connProfile)

	fmt.Printf("%s",execute("clear"))
	fmt.Println("						Os dados informados foram: \n\n" +

			"					Interface de rede: " + connProfile.wifiInterface + "\n\n" +

			"					Tipo de Conexao: " + connProfile.connectionType  + "\n\n" +

			"					Seguranca da rede: "+ connProfile.wifiSecurityType +"\n\n" +

			"					Nome da Rede Wi-fi: "+ connProfile.essid + "\n\n" +

			"					Modo de aquisicao IP: "+ connProfile.ipMode + "\n\n" +

			"					Senha da Rede: " + connProfile.wifiPassword + "\n\n" +

			"					Rede Oculta: " + connProfile.hidden + "\n") 
	




//	execute("loadkeys br-abnt2")
//	connProfile.writeWifiConfigToFile(".", "teste.txt")


//	netctl start wifiInterface  //substitui o wifi-menu
//	timedatectl set-ntp true


//	fmt.Println("Informe em qual dispositivo você deseja criar o particionamento")
//	fmt.Println(detectPartitionTable())
//	fmt.Scanf("%s", &)

//	fmt.Println("Seu computador tem suporte a EFI?")
//	fmt.Prinln("[yes no]")
//	fmt.Scanf("%s",&)

//	fmt.Println("Você deseja utilizar todo o espaço da partição para o sistema?")
//	fmt.Println("[yes]")
//	fmt.Scanf("%s",&)
	//createPartitionTable(wifiInterface)

	//parted

//	fmt.Println("O sistema será instalado em um pendrive ou ssd?")
//	fmt.Println("[yes]")
//	fmt.Scanf("%s",&)


//	fmt.Println("Você deseja instalar então o sistema de arquivos f2fs, que permite um melhor aproveitamento para esse tipo de disposivo?")
//	fmt.Println("[yes]")
//	fmt.Scanf("%s",&)
	//mkfs.f2fs

	//mkdir -p /mnt/boot
	//mount ????

	//pacstrap -i /mnt base base-devel
	//genfstab -U /mnt > /mnt/etc/fstab
	//arch-chroot /mnt /bin/bash
	//locale-gen



//	fmt.Println("Qual língua você deseja que o sistema possua?")
//	fmt.Scanf("%s",&)

	//configurar /etc/locale.conf LANG=pt_BR.???

//	fmt.Println("Existe um padrão de teclado definido para sua linguagem. Você deseja informar alguma configuração fora do padrão?")
//	fmt.Println("[no]")
//	fmt.Scanf("%s",&)
	//KEYMAP=br.abnt2  -- colocar no /etc/vconsole


	//tzselect
	//ln -sf /usr/share/zoneinfo/Zone???/Subzone??? /etc/localtime


	//hwclock --systohc --utc
	//mkinitcpio -p linux
	
 	
}
