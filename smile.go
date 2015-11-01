// smile


package main

import (
	"fmt"
	"os/exec"
	"strings"
	"io/ioutil"
)

const version ="0.0.0"

func check(e error) {
	if e != nil {
		panic(e)
	}	
}
func boolYesNo(value bool) string{
	if value == true{
		return "yes"
	}else{
		return "no"
	}		
}

func execute(cmd string) {

	cmdLine := strings.Fields(cmd)
	command := cmdLine[0]
	parameters := cmdLine[1:len(cmdLine)]

	out, err := exec.Command(command, parameters...).Output()

	check(err)

	fmt.Printf("%s", out)

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

func createWifiConfig(connProfile *ConnectionProfile) {

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

		     descHidden + connProfile.hidden

	err:=ioutil.WriteFile("teste.txt",data,0777)
	check(err)


}

func createPartitionTable(device string){


	fmt.Printf("%s",device)


}

func main() {

	connProfile:={
//	execute("loadkeys br-abnt2")
	createWifiConfig(wifiUser,wifiPass,hidden,wifiInterface)
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
