package cli

import "fmt"

type Command struct {
	Action string
}

func ParseArgs(args []string) (*Command, error) {
	if len(args) == 0 {
		return nil, fmt.Errorf("no arguments provided")
	}
	
	switch args[0] {
	case "--success":
		return &Command{Action: "success"}, nil
	case "--failure":
		return &Command{Action: "failure"}, nil
	case "--init-scenes":
		return &Command{Action: "init-scenes"}, nil
	default:
		return nil, fmt.Errorf("invalid argument: %s", args[0])
	}
}