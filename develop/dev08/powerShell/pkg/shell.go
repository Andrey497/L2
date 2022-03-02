package pkg

type shell struct {
	command    string
	parameters string
}

func NewShell(command, parameter string) *shell {
	return &shell{command: command, parameters: parameter}
}
func (sh *shell) Execute() error {
	return sh.switchCommand()
}
func (sh *shell) switchCommand() error {
	switch sh.command {
	case "cd":
		return CommandCd(sh.parameters)
	case "pwd":
		return CommandPwd()
	case "echo":
		return CommandEcho(sh.parameters)
	case "kil":
		return CommandKill(sh.parameters)
	case "ps":
		return CommandPs()
	}

	return nil
}
