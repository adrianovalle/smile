package main

import (
	"os"
	"os/exec"
	"fmt"
)

func Start(args ...string) (p *os.Process, err error) {
	if args[0], err = exec.LookPath(args[0]); err == nil {
		var procAttr os.ProcAttr
		procAttr.Files = []*os.File{os.Stdin,
			os.Stdout, os.Stderr}
		p, err := os.StartProcess(args[0], args, &procAttr)
		if err == nil {
			return p, nil
		}
	}
	return nil, err
}

func main() {
	
	if proc, err := Start("bash", "-c","arch-chroot '/mnt' '/bin/bash' -c 'mkdir teste ;ls > teste.txt'"); err== nil{

	
	
	//if proc, err := Start("bash", "-c", "mkdir hello" +
	//			";" + "mkdir badabes"); err == nil {
		fmt.Println("Feliz Natal, ho ho ho")
		
		
		proc.Wait()
	}
//	if proc, err := Start("zsh"); err == nil {
//		proc.Wait()
//	}
}
