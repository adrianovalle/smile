// smile
package main

import (
	"fmt"
	"os/exec"
	"strings"
	"io/ioutil"
)

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

	const description = "Description=Configuration file created by smile 0.0.0"
	const descInterface = "Interface="
	const descConnectionType = "Connection="
	const descWifiSecurityType = "Security="
	const descESSID = "ESSID="
	const descIPMode = "IP="
	const descWifiPass = "Key="
	const descHidden = "Hidden="


	data:=[]byte(wifiInterface


		     wifiUser  +"\n"+
		     wifiPass  +"\n"+
		     boolYesNo(hidden)    +"\n"+
		     wifiInterface  +"\n")

	err:=ioutil.WriteFile("teste.txt",data,0777)
	check(err)

	fmt.Printf("%s\n", wifiUser)
	fmt.Printf("%s\n", wifiPass)
	fmt.Printf("%v\n", hidden)
	fmt.Printf("%s\n", wifiInterface)



}

func createPartitionTable(device string){


	fmt.Printf("%s",device)


}

func main() {

	wifiUser := "usuario"
	wifiPass := "senha"
	hidden := true
	wifiInterface := "wlp2s0"

	execute("loadkeys br-abnt2")
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
