// smile


package main

import (
	"fmt"
	"os/exec"
	"strings"
	"regexp"


)


const version ="0.22.0"

func check(e error) {
	if e != nil {
		panic(e)
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

func main() {

	a:="enp3s0 wlp2s0 amor casa"
	r,_:=regexp.Compile("[a-z]{3}[0-9][a-z][0-9]")
	b:=r.FindAllString(a,-1)
	//lsblk
	//createPartitionTable(wifiInterface)
	//parted
	//mkfs
	//mkdir -p /mnt/boot
	//mount ????
}

