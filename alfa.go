package main

import (
//	"fmt"
//	"io/ioutil"
//	"os"
	"os/exec"
//	"regexp"
	"strings"
//	"time"
)

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
//	if verbose == true {
//		fmt.Printf("%s", out)
//	}
	return out
}
func main(){

	_ = execute ("arch-chroot /mnt /bin/bash -c 'mkdir hadouken'")


}
