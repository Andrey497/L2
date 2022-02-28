package pkg

type shell struct {
	command    string
	parameters string
}

func NewShell(command, parameter string) *shell {
	return &shell{command: command, parameters: parameter}
}
func (sh *shell) Execute() error {
	sh.switchCommand()
	return nil
}
func (sh *shell) switchCommand() error {
	switch sh.command {
	case "cd":

	case "pwd":
		return CommandPwd()
	case "echo":
		return CommandEcho(sh.parameters)
	case "kil":
	case "ps":

	}

	return nil
}
