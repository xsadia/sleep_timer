package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func main() {

	osArgs := os.Args[1:]

	if len(osArgs) < 1 {
		log.Fatal("Not enough arguments")
	}

	pid, err := getPid("firefox")

	if err != nil {
		log.Fatal("Something went wrong")
	}

	v, _ := strconv.Atoi(osArgs[0])

	time.Sleep(time.Duration(v) * time.Minute)

	if err = kill(pid); err != nil {
		log.Fatalf("Failed to kill process of pid: %s", pid)
	}

	fmt.Printf("Successfully killed process of pid: %s", pid)

	if err = suspend(); err != nil {
		log.Fatal("Failed to suspend")
	}

	os.Exit(0)
}

func getPid(process string) ([]byte, error) {
	pid, err := exec.Command("pgrep", process).Output()

	if err != nil {
		return []byte{}, err
	}

	return pid, nil
}

func kill(pid []byte) error {
	_, err := exec.Command("kill", "-SIGTERM", strings.Trim(string(pid), "\n")).Output()

	return err
}

func suspend() error {
	_, err := exec.Command("systemctl", "suspend").Output()

	return err
}
