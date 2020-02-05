package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("/bin/bash", "-c", "open /Applications/minecraft.app/Contents/MacOS/HMCL.jar")
	buf, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(buf))
}
