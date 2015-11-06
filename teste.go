// smile

package main

import (
	"fmt"
	"os/exec"
//	"strings"
	"regexp"

)

const version = "0.22.0"

func check(e error) {
	if e != nil {
		panic(e)
	}
}


func main() {

	a,err:= exec.Command("ip", "link").Output()
	d,_:=exec.Command("clear").Output()

	b:=string(a)
	check(err)	
	fmt.Printf("%s",d)
	//	r,_:=regexp.Compile("[a-z]{3}[0-9][a-z][0-9]")
	//  a:="amor casa wlan0 wlp2s0":x

		r,_:=regexp.Compile("[e-w]{1}[a-z]{3}[0-9]")
		c:=r.FindAllString(b,-1)
	fmt.Println(c)
	//lsblk
	//createPartitionTable(wifiInterface)
	//parted
	//mkfs
	//mkdir -p /mnt/boot
	//mount ????
}
