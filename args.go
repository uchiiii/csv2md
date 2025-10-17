package main

import (
	"fmt"
)

const (
	ERROR_NO_FILENAME = "Error: at least one file name is required."
)

type Args struct {
	Files      []string
	Delim      string
	Pad        int
	OutputFile string
}

var Delims = []string{",", ";"}

func (args *Args) ValidateAll() error {
	if err := args.validateFiles(); err != nil {
		return err
	}
	if err := args.validateDelim(); err != nil {
		return err
	}
	if err := args.validatePad(); err != nil {
		return err
	}

	return nil
}

func (args *Args) validateFiles() error {
	if len(args.Files) < 1 {
		return fmt.Errorf("%s", ERROR_NO_FILENAME)
	}
	return nil
}

func (args *Args) validateDelim() error {
	if !contains(Delims, args.Delim) {
		return fmt.Errorf("Error: delimiter should be one of %+q, but got %q", Delims, args.Delim)
	}
	return nil
}

func (args *Args) validatePad() error {
	if args.Pad < 0 {
		return fmt.Errorf("Error: padding should be positive integer, but got %d", args.Pad)
	}
	return nil
}

func contains(s []string, a string) bool {
	for _, v := range s {
		if v == a {
			return true
		}
	}
	return false
}
