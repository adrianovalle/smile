// smile

package main

import (
	"fmt"
	"os/exec"
	"os"

)

const version = "0.22.0"

func check(e error) {
	if e != nil {
		panic(e)
	}
}


func execute(cmdLine string) []byte{






	return out

}







func main() {

	a,err:= exec.Command("ip", "link").Output()
	d,_:=exec.Command("clear").Output()

	b:=string(a)
	check(err)	
	fmt.Printf("%s",d)
}
