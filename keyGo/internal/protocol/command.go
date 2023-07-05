package protocol

import "strings"

type Command struct {
	Name string
	Args []string
}

func ParseCommand(text string) Command {

	words := strings.Fields(text)

	name := words[0]
	args := words[1:]

	return Command{Name: name, Args: args}

}
