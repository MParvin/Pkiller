package tools

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"
)

func KillProc(pid int) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Are you sure you want to kill process with PID", pid, "? (y/N)")
	choice, _ := reader.ReadString('\n')
	choice = strings.ToLower(choice)
	if choice == "y\n" {
		syscall.Kill(pid, syscall.SIGTERM)
	}

}
