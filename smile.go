// smile


package main

import (
	"fmt"
	"os/exec"
	"strings"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
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
	hidden		 bool


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

		     descHidden + strconv.FormatBool(connProfile.hidden))


	err:=ioutil.WriteFile(destinationFolder + "/" + nameProfile,data,0777)
	check(err)


}
func copyFile(originalFileWithPath string, destinyFileWithPath string){

	err := os.Rename(originalFileWithPath, destinyFileWithPath)

	check(err)

}
func detectNetwork() []string{


	cmdOut:=execute("ip link")
	
//	r,_:=regexp.Compile("[a-z]{3}[0-9][a-z][0-9]")


	r,_:=regexp.Compile("[ew]{1}[a-z]{3}[0-9]")
	
	regex:=r.FindAllString(string(cmdOut), -1)
	return regex
}
func createPartitionTable(device string){


	fmt.Printf("%s",device)


}

func main() {
var wifiInterface, connectionType, wifiSecurityType, essid,ipMode, wifiPassword string
var hidden bool


	verbose=false

	_=execute("clear")

	fmt.Printf("Bom dia! Informe sua interface de rede \n")

	fmt.Println(detectNetwork())
	
	fmt.Scanf("%s\n",&wifiInterface)

	fmt.Printf("Informe o tipo de conexão \n")
	fmt.Scanf("%s\n" ,&connectionType)

	fmt.Printf("Informe a segurança wi-fi \n")
	fmt.Scanf("%s\n", &wifiSecurityType)

	fmt.Printf("Informe a identificação da rede \n")
	fmt.Scanf("%s\n" , &essid)

	fmt.Printf("Informe o modo de Ip \n")
	fmt.Scanf("%s\n", &ipMode)

	fmt.Printf("Informe a senha da rede \n")
	fmt.Scanf("%s\n", &wifiPassword)

	fmt.Printf("A rede está oculta? \n")
	fmt.Scanf("%t\n", &hidden)

	connProfile := ConnectionProfile{wifiInterface, connectionType, wifiSecurityType, essid, ipMode, wifiPassword, hidden}
//	execute("loadkeys br-abnt2")
	connProfile.writeWifiConfigToFile(".", "teste.txt")
	//netctl start wifiInterface  //substitui o wifi-menu
	//timedatectl set-ntp true
	//lsblk
	//createPartitionTable(wifiInterface)
	//parted
	//mkfs
	//mkdir -p /mnt/boot
	//mount ????
	//pacstrap -i /mnt base base-devel
	//genfstab -U /mnt > /mnt/etc/fstab
	//arch-chroot /mnt /bin/bash
	//locale-gen
	//configurar /etc/locale.conf LANG=pt_BR.???
	//KEYMAP=br.abnt2  -- colocar no /etc/vconsole
	//tzselect
	//ln -sf /usr/share/zoneinfo/Zone???/Subzone??? /etc/localtime
	//hwclock --systohc --utc
	//mkinitcpio -p linux
	
 	
}
