package main

import (
//	"fmt"
//	"io/ioutil"
//	"os"
	"os/exec"
//	"regexp"
//	"strings"
//	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func execute()[]byte {


	out, err := exec.Command("bash", "-c","arch-chroot '/mnt' '/bin/bash' -c 'mkdir teste ;ls > teste.txt'").Output()

	check(err)
//	if verbose == true {
//		fmt.Printf("%s", out)
//	}
	return out
}
func main(){

	_ = execute()


}
