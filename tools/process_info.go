package tools

import (
	"bufio"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func IsProcessRunning(pid int) bool {
	_, err := os.Stat("/proc/" + strconv.Itoa(pid))
	return err == nil
}

func GetProcInfo(pid int) (string, string, error) {
	pCmd := getCmd(pid)
	uid := getUid(pid)
	uName := getUser(uid)
	return pCmd, uName, nil
}

func getCmd(pid int) string {
	filePath := "/proc/" + strconv.Itoa(pid) + "/cmdline"
	cmdLine, _ := ioutil.ReadFile(filePath)
	return string(cmdLine)
}

func getUid(pid int) string {
	filePath := "/proc/" + strconv.Itoa(pid) + "/status"
	status, _ := ioutil.ReadFile(filePath)
	statusLines := strings.Split(string(status), "\n")
	var uidLine string
	for _, line := range statusLines {
		if strings.HasPrefix(line, "Uid:") {
			uidLine = line
			break
		}
	}
	uid := strings.Split(uidLine, "\t")[1]
	return uid
}

func getUser(uid string) string {
	file, _ := os.Open("/etc/passwd")
	defer file.Close()

	scanr := bufio.NewScanner(file)
	for scanr.Scan() {
		line := scanr.Text()
		parts := strings.Split(line, ":")
		if parts[2] == uid {
			return parts[0]
		}
	}
	return ""
}
