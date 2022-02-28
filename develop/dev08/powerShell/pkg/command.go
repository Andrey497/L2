package pkg

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func command(command, parameter string) error {
	var cmd *exec.Cmd
	if strings.TrimSpace(parameter) != "" {
		cmd = exec.Command(command, parameter)
	} else {
		cmd = exec.Command(command, parameter)
	}
	stdout, err := cmd.StdoutPipe()

	if err != nil {
		return err
	}

	if err := cmd.Start(); err != nil {
		return err
	}

	data, err := ioutil.ReadAll(stdout)

	if err != nil {
		return err
	}

	if err := cmd.Wait(); err != nil {
		return err
	}
	fmt.Print(string(data))
	return nil
}
func CommandCd(parameter string) error {
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}
	err = os.Chdir(parameter)
	if err != nil {
		return err
	}
	fmt.Println("pre director:" + pwd)
	pwd, err = os.Getwd()
	if err != nil {
		return err
	}
	fmt.Println("current director:" + pwd)
	return command("cd", parameter)

}
func CommandPwd() error {
	return command("pwd", "")
	//v2
	//pwd, err := os.Getwd()
	//if err != nil {
	//	return err
	//}
	//fmt.Println(pwd)
}
func CommandEcho(parameter string) error {
	return command("echo", parameter)
}
func CommandKill(parameter string) error {
	if strings.TrimSpace(parameter) == "" {
		fmt.Println("input id kill process")
		fmt.Scan(&parameter)
		_, err := strconv.Atoi(parameter)
		if err != nil {
			return err
		}
		return command("kill", parameter)
	}
	return nil
}
func CommandPs() error {
	return command("ps", "")
}
