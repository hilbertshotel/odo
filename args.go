package main

import (
	"errors"
	"os"
)

func parseArgs() (string, string, error) {
	args := os.Args[1:]
	length := len(args)
	cmdErr := errors.New("invalid command")
	argsErr := errors.New("too many arguments")

	switch length {
	case 0:
		return "info", "", nil
	case 1:
		cmd := args[0]
		if cmd == "play" {
			return cmd, "", nil
		} else {
			return "", "", cmdErr
		}
	case 2:
		cmd, name := args[0], args[1]
		if cmd == "new" {
			return cmd, name, nil
		} else if cmd == "build" {
			return cmd, name, nil
		} else {
			return "", "", cmdErr
		}
	default:
		return "", "", argsErr
	}
}
