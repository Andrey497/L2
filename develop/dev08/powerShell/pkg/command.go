package pkg

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

func CommandCd(parameter string) {

}
func CommandPwd() error {
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}
	fmt.Println(pwd)
	return nil
}
func CommandEcho(parameter string) error {
	cmd := exec.Command("echo", parameter)
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
func CommandKill(parameter string) {

}
func CommandPs(parameter string) {

}
