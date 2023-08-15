package main

import "os/exec"

//多进程适用于计算密集型，消耗cpu
//多线程适用于io密集型
func main() {
	//cmd := exec.Command("notepad")
	cmd := exec.Command("taskkill", "/f", "/im", "notepad.exe")
	cmd.Run()
}
