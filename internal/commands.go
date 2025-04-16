package internal

type CliCommand struct {
	Name        string
	Description string
	Callback    func(c *Config, args []string) error
}

var SupportedCommands = map[string]CliCommand{}

func AddCliCommand(name string, description string, callback func(c *Config, args []string) error) {
	SupportedCommands[name] = CliCommand{name, description, callback}
}

func GetCliCommand(name string) (CliCommand, bool) {
	cmd, exists := SupportedCommands[name]
	return cmd, exists
}
